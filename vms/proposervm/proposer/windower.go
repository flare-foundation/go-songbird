// (c) 2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposer

import (
	"github.com/flare-foundation/flare/vms/rpcchainvm"
	"sort"
	"time"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/validators"
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
		chainHeight,
		pChainHeight uint64,
		validatorID ids.ShortID,
		hash ids.ID,
	) (time.Duration, error)
	GetValidators(hash []byte) (map[string]float64, error)
}

// windower interfaces with P-Chain and it is responsible for calculating the
// delay for the block submission window of a given validator
type windower struct {
	state       validators.State
	subnetID    ids.ID
	chainSource uint64
	sampler     sampler.WeightedWithoutReplacement
	//vmValidator  *validatorvm.ValidatorVM
	valClient *rpcchainvm.ValidatorsClient
	//valInterface *block.ValidatorVMInterface
}

func New(state validators.State, subnetID, chainID ids.ID) Windower {
	w := wrappers.Packer{Bytes: chainID[:]}
	return &windower{
		state:       state,
		subnetID:    subnetID,
		chainSource: w.UnpackLong(),
		sampler:     sampler.NewDeterministicWeightedWithoutReplacement(),
		valClient:   rpcchainvm.GlobalValidatorClient,
	}
}

func (w *windower) Delay(chainHeight, pChainHeight uint64, validatorID ids.ShortID, hash ids.ID) (time.Duration, error) {
	validatorsMapNew, err := w.valClient.GetValidators(hash)
	if validatorID == ids.ShortEmpty {
		return MaxDelay, nil
	}

	// get the validator set by the p-chain height
	validatorsMap, err := w.state.GetValidatorSet(pChainHeight, w.subnetID)
	if err != nil {
		return 0, err
	}

	// convert the map of validators to a slice
	validators := make(validatorsSlice, 0, len(validatorsMap))
	weight := uint64(0)
	for k, v := range validatorsMap {
		validators = append(validators, validatorData{
			id:     k,
			weight: v,
		})
		newWeight, err := math.Add64(weight, v)
		if err != nil {
			return 0, err
		}
		weight = newWeight
	}

	// New validators from coreth
	validators = nil
	for id, u := range validatorsMapNew {
		//longID, err := ids.ToID(id)
		//shortID, err := ids.ToShortID(id)
		if err != nil {
			continue
		}
		validators = append(validators, validatorData{
			id:        id, //todo figure out why shortID is used
			weightNew: u,
		})
	}

	// canonically sort validators
	// Note: validators are sorted by ID, sorting by weight would not create a
	// canonically sorted list
	sort.Sort(validators)

	// convert the slice of validators to a slice of weights
	validatorWeights := make([]uint64, len(validators))
	for i, v := range validators {
		validatorWeights[i] = v.weight
	}

	if err := w.sampler.Initialize(validatorWeights); err != nil {
		return 0, err
	}

	numToSample := MaxWindows
	if weight < uint64(numToSample) {
		numToSample = int(weight)
	}

	seed := chainHeight ^ w.chainSource
	w.sampler.Seed(int64(seed))

	indices, err := w.sampler.Sample(numToSample)
	if err != nil {
		return 0, err
	}

	delay := time.Duration(0)
	for _, index := range indices {
		nodeID := validators[index].id
		if nodeID == validatorID {
			return delay, nil
		}
		delay += WindowDuration
	}
	return delay, nil
}

func (w *windower) GetValidators(hash []byte) (map[string]float64, error) {
	var id [32]byte
	copy(id[:], hash)

	m, err := w.valClient.GetValidators(id)
	if err != nil {
		return nil, err
	}
	return convertShortIdmapToStringMap(m), nil

}

func convertShortIdmapToStringMap(m map[ids.ShortID]float64) map[string]float64 {
	retM := make(map[string]float64)
	for key, val := range m {
		retM[key.String()] = val
	}
	return retM
}