// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposer

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/validators"
)

func TestWindowerNoValidators(t *testing.T) {
	assert := assert.New(t)

	chainID := ids.GenerateTestID()
	nodeID := ids.GenerateTestShortID()

	noValidators := validators.NewSet()

	retriever := &validators.TestRetriever{
		GetValidatorsByBlockIDF: func(blockID ids.ID) (validators.Set, error) {
			return noValidators, nil
		},
	}

	w := New(retriever, chainID)

	delay, err := w.Delay(1, ids.ID{}, nodeID)
	assert.NoError(err)
	assert.EqualValues(0, delay)
}

func TestWindowerRepeatedValidator(t *testing.T) {
	assert := assert.New(t)

	chainID := ids.GenerateTestID()
	validatorID := ids.GenerateTestShortID()
	nonValidatorID := ids.GenerateTestShortID()

	retriever := &validators.TestRetriever{
		GetValidatorsByBlockIDF: func(blockID ids.ID) (validators.Set, error) {
			s := validators.NewSet() //todo use the validatorID in NewSet and NOT use the nonValidatorID
			s.AddWeight(validatorID, 10)
			return s, nil
		},
	}

	w := New(retriever, chainID)

	validatorDelay, err := w.Delay(1, ids.ID{}, validatorID)
	assert.NoError(err)
	assert.EqualValues(0, validatorDelay)

	nonValidatorDelay, err := w.Delay(1, ids.ID{}, nonValidatorID)
	assert.NoError(err)
	assert.EqualValues(MaxDelay, nonValidatorDelay)
}

func TestWindowerChangeByHeight(t *testing.T) {
	assert := assert.New(t)

	chainID := ids.ID{0, 2}
	validatorIDs := make([]ids.ShortID, MaxWindows)
	for i := range validatorIDs {
		validatorIDs[i] = ids.ShortID{byte(i + 1)}
	}

	retriever := &validators.TestRetriever{
		GetValidatorsByBlockIDF: func(blockID ids.ID) (validators.Set, error) {
			s := validators.NewSet()
			for _, id := range validatorIDs {
				s.AddWeight(id, 1)
			}
			return s, nil
		},
	}

	w := New(retriever, chainID)

	expectedDelays1 := []time.Duration{
		2 * WindowDuration,
		5 * WindowDuration,
		3 * WindowDuration,
		4 * WindowDuration,
		0 * WindowDuration,
		1 * WindowDuration,
	}
	for i, expectedDelay := range expectedDelays1 {
		vdrID := validatorIDs[i]
		fmt.Println(vdrID)
		validatorDelay, err := w.Delay(1, ids.ID{}, vdrID)
		assert.NoError(err)
		assert.EqualValues(expectedDelay, validatorDelay)
	}

	expectedDelays2 := []time.Duration{
		5 * WindowDuration,
		1 * WindowDuration,
		3 * WindowDuration,
		4 * WindowDuration,
		0 * WindowDuration,
		2 * WindowDuration,
	}
	for i, expectedDelay := range expectedDelays2 {
		vdrID := validatorIDs[i]
		validatorDelay, err := w.Delay(2, ids.ID{}, vdrID)
		assert.NoError(err)
		assert.EqualValues(expectedDelay, validatorDelay)
	}
}

func TestWindowerChangeByChain(t *testing.T) {
	assert := assert.New(t)

	rand.Seed(0)
	chainID0 := ids.ID{}
	_, _ = rand.Read(chainID0[:]) // #nosec G404
	chainID1 := ids.ID{}
	_, _ = rand.Read(chainID1[:]) // #nosec G404

	validatorIDs := make([]ids.ShortID, MaxWindows)
	for i := range validatorIDs {
		validatorIDs[i] = ids.ShortID{byte(i + 1)}
	}

	retriever := &validators.TestRetriever{
		GetValidatorsByBlockIDF: func(blockID ids.ID) (validators.Set, error) {
			s := validators.NewSet()
			for _, id := range validatorIDs {
				s.AddWeight(id, 1)
			}
			return s, nil
		},
	}

	w0 := New(retriever, chainID0)
	w1 := New(retriever, chainID1)

	expectedDelays0 := []time.Duration{
		5 * WindowDuration,
		2 * WindowDuration,
		0 * WindowDuration,
		3 * WindowDuration,
		1 * WindowDuration,
		4 * WindowDuration,
	}
	for i, expectedDelay := range expectedDelays0 {
		vdrID := validatorIDs[i]
		validatorDelay, err := w0.Delay(1, ids.ID{}, vdrID)
		assert.NoError(err)
		assert.EqualValues(expectedDelay, validatorDelay)
	}

	expectedDelays1 := []time.Duration{
		0 * WindowDuration,
		1 * WindowDuration,
		4 * WindowDuration,
		5 * WindowDuration,
		3 * WindowDuration,
		2 * WindowDuration,
	}
	for i, expectedDelay := range expectedDelays1 {
		vdrID := validatorIDs[i]
		validatorDelay, err := w1.Delay(1, ids.ID{}, vdrID)
		assert.NoError(err)
		assert.EqualValues(expectedDelay, validatorDelay)
	}
}
