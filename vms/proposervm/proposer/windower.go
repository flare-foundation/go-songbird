// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposer

import (
	"fmt"
	"sort"
	"time"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/validation"
	"github.com/flare-foundation/flare/utils/math"
	"github.com/flare-foundation/flare/utils/sampler"
	"github.com/flare-foundation/flare/utils/wrappers"
)

// Proposer list constants
const (
	MaxWindows     = 6
	WindowDuration = 5 * time.Second
	MaxDelay       = MaxWindows * WindowDuration
)

var _ Windower = &windower{}

type Windower interface {
	Delay(
		height uint64,
		parentID ids.ID,
		validatorID ids.ShortID,
	) (time.Duration, error)
}

// windower interfaces with P-Chain and it is responsible for calculating the
// delay for the block submission window of a given validator
type windower struct {
	retriever validation.Retriever
	nonce     uint64
	sampler   sampler.WeightedWithoutReplacement
}

func New(retriever validation.Retriever, chainID ids.ID) Windower {
	w := wrappers.Packer{Bytes: chainID[:]}
	return &windower{
		retriever: retriever,
		nonce:     w.UnpackLong(),
		sampler:   sampler.NewDeterministicWeightedWithoutReplacement(),
	}
}

func (w *windower) Delay(height uint64, parentID ids.ID, validatorID ids.ShortID) (time.Duration, error) {

	if validatorID == ids.ShortEmpty {
		return MaxDelay, nil
	}

	validators, err := w.retriever.GetValidators(parentID)
	if err != nil {
		return 0, fmt.Errorf("could not get validators for windowing: %w", err)
	}

	// canonically sort validators
	// Note: validators are sorted by ID, sorting by weight would not create a
	// canonically sorted list
	validatorList := validators.List()
	sort.Sort(sortByID(validatorList))

	// Then, create slices of weights and IDs for sampling.
	totalWeight := uint64(0)
	validatorIDs := make([]ids.ShortID, 0, len(validatorList))
	weights := make([]uint64, 0, len(validatorList))
	for _, validator := range validatorList {
		totalWeight, err = math.Add64(totalWeight, validator.Weight())
		if err != nil {
			return 0, err
		}
		validatorIDs = append(validatorIDs, validator.ID())
		weights = append(weights, validator.Weight())
	}

	if err := w.sampler.Initialize(weights); err != nil {
		return 0, err
	}

	numToSample := MaxWindows
	if totalWeight < uint64(numToSample) {
		numToSample = int(totalWeight)
	}

	seed := height ^ w.nonce
	w.sampler.Seed(int64(seed))

	indices, err := w.sampler.Sample(numToSample)
	if err != nil {
		return 0, err
	}

	delay := time.Duration(0)
	for _, index := range indices {
		nodeID := validatorIDs[index]
		if nodeID == validatorID {
			return delay, nil
		}
		delay += WindowDuration
	}

	return delay, nil
}
