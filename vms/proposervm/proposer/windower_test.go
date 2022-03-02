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
)

// FIXME: fix these tests after validator changes

func TestWindowerNoValidators(t *testing.T) {
	assert := assert.New(t)

	chainID := ids.GenerateTestID()
	nodeID := ids.GenerateTestShortID()
	parentID := ids.GenerateTestID()

	w := New(nil, chainID)

	delay, err := w.Delay(1, parentID, nodeID)
	assert.NoError(err)
	assert.EqualValues(0, delay)
}

func TestWindowerRepeatedValidator(t *testing.T) {
	assert := assert.New(t)

	chainID := ids.GenerateTestID()
	parentID := ids.GenerateTestID()
	validatorID := ids.GenerateTestShortID()
	nonValidatorID := ids.GenerateTestShortID()

	w := New(nil, chainID)

	// FIXME: double check parent ID
	validatorDelay, err := w.Delay(1, parentID, validatorID)
	assert.NoError(err)
	assert.EqualValues(0, validatorDelay)

	// FIXME: double check parent ID
	nonValidatorDelay, err := w.Delay(1, parentID, nonValidatorID)
	assert.NoError(err)
	assert.EqualValues(MaxDelay, nonValidatorDelay)
}

func TestWindowerChangeByHeight(t *testing.T) {
	assert := assert.New(t)

	chainID := ids.ID{0, 2}
	parentID := ids.ID{0, 3}
	validatorIDs := make([]ids.ShortID, MaxWindows)
	for i := range validatorIDs {
		validatorIDs[i] = ids.ShortID{byte(i + 1)}
	}

	w := New(nil, chainID)

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
		// FIXME: double check parent ID
		validatorDelay, err := w.Delay(1, parentID, vdrID)
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
		// FIXME: double check parent ID
		validatorDelay, err := w.Delay(2, parentID, vdrID)
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

	parentID := ids.ID{}
	_, _ = rand.Read(parentID[:])

	validatorIDs := make([]ids.ShortID, MaxWindows)
	for i := range validatorIDs {
		validatorIDs[i] = ids.ShortID{byte(i + 1)}
	}

	w0 := New(nil, chainID0)
	w1 := New(nil, chainID1)

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
		// FIXME: double check parent ID
		validatorDelay, err := w0.Delay(1, parentID, vdrID)
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
		// FIXME: double check parent ID
		validatorDelay, err := w1.Delay(1, parentID, vdrID)
		assert.NoError(err)
		assert.EqualValues(expectedDelay, validatorDelay)
	}
}
