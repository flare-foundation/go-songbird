// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package benchlist

import (
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/stretchr/testify/assert"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/validation"
	"github.com/flare-foundation/flare/utils/logging"
	"github.com/flare-foundation/flare/utils/wrappers"
)

var minimumFailingDuration = 5 * time.Minute

// Test that validators are properly added to the bench
func TestBenchlistAdd(t *testing.T) {
	validators := validation.NewSet()
	validator1 := validation.GenerateRandomValidator(50)
	validator2 := validation.GenerateRandomValidator(50)
	validator3 := validation.GenerateRandomValidator(50)
	validator4 := validation.GenerateRandomValidator(50)
	validator5 := validation.GenerateRandomValidator(50)

	errs := wrappers.Errs{}
	errs.Add(
		validators.AddWeight(validator1.ID(), validator1.Weight()),
		validators.AddWeight(validator2.ID(), validator2.Weight()),
		validators.AddWeight(validator3.ID(), validator3.Weight()),
		validators.AddWeight(validator4.ID(), validator4.Weight()),
		validators.AddWeight(validator5.ID(), validator5.Weight()),
	)
	if errs.Errored() {
		t.Fatal(errs.Err)
	}

	benchable := &TestBenchable{T: t}
	benchable.Default(true)

	threshold := 3
	duration := time.Minute
	maxPortion := 0.5
	benchIntf, err := NewBenchlist(
		ids.Empty,
		logging.NoLog{},
		benchable,
		validators,
		threshold,
		minimumFailingDuration,
		duration,
		maxPortion,
		prometheus.NewRegistry(),
	)
	if err != nil {
		t.Fatal(err)
	}
	b := benchIntf.(*benchlist)
	defer b.timer.Stop()
	now := time.Now()
	b.clock.Set(now)

	// Nobody should be benched at the start
	b.lock.Lock()
	assert.False(t, b.isBenched(validator1.ID()))
	assert.False(t, b.isBenched(validator2.ID()))
	assert.False(t, b.isBenched(validator3.ID()))
	assert.False(t, b.isBenched(validator4.ID()))
	assert.False(t, b.isBenched(validator5.ID()))
	assert.Len(t, b.failureStreaks, 0)
	assert.Equal(t, b.benchedQueue.Len(), 0)
	assert.Equal(t, b.benchlistSet.Len(), 0)
	b.lock.Unlock()

	// Register [threshold - 1] failures in a row for validator1
	for i := 0; i < threshold-1; i++ {
		b.RegisterFailure(validator1.ID())
	}

	// Still shouldn't be benched due to not enough consecutive failure
	assert.False(t, b.isBenched(validator1.ID()))
	assert.Equal(t, b.benchedQueue.Len(), 0)
	assert.Equal(t, b.benchlistSet.Len(), 0)
	assert.Len(t, b.failureStreaks, 1)
	fs := b.failureStreaks[validator1.ID()]
	assert.Equal(t, threshold-1, fs.consecutive)
	assert.True(t, fs.firstFailure.Equal(now))

	// Register another failure
	b.RegisterFailure(validator1.ID())

	// Still shouldn't be benched because not enough time (any in this case)
	// has passed since the first failure
	b.lock.Lock()
	assert.False(t, b.isBenched(validator1.ID()))
	assert.Equal(t, b.benchedQueue.Len(), 0)
	assert.Equal(t, b.benchlistSet.Len(), 0)
	b.lock.Unlock()

	// Move the time up
	now = now.Add(minimumFailingDuration).Add(time.Second)
	b.lock.Lock()
	b.clock.Set(now)

	benched := false
	benchable.BenchedF = func(ids.ID, ids.ShortID) {
		benched = true
	}
	b.lock.Unlock()

	// Register another failure
	b.RegisterFailure(validator1.ID())

	// Now this validator should be benched
	b.lock.Lock()
	assert.True(t, b.isBenched(validator1.ID()))
	assert.Equal(t, b.benchedQueue.Len(), 1)
	assert.Equal(t, b.benchlistSet.Len(), 1)

	next := b.benchedQueue[0]
	assert.Equal(t, validator1.ID(), next.validatorID)
	assert.True(t, !next.benchedUntil.After(now.Add(duration)))
	assert.True(t, !next.benchedUntil.Before(now.Add(duration/2)))
	assert.Len(t, b.failureStreaks, 0)
	assert.True(t, benched)
	benchable.BenchedF = nil
	b.lock.Unlock()

	// Give another validator [threshold-1] failures
	for i := 0; i < threshold-1; i++ {
		b.RegisterFailure(validator2.ID())
	}

	// Advance the time
	b.lock.Lock()
	now = now.Add(minimumFailingDuration)
	b.lock.Unlock()

	// Register another failure
	b.RegisterResponse(validator2.ID())

	// validator2 shouldn't be benched
	// The response should have cleared its consecutive failures
	b.lock.Lock()
	assert.True(t, b.isBenched(validator1.ID()))
	assert.False(t, b.isBenched(validator2.ID()))
	assert.Equal(t, b.benchedQueue.Len(), 1)
	assert.Equal(t, b.benchlistSet.Len(), 1)
	assert.Len(t, b.failureStreaks, 0)
	b.lock.Unlock()

	// Register another failure for validator1, who is benched
	b.RegisterFailure(validator1.ID())

	// A failure for an already benched validator should not count against it
	b.lock.Lock()
	assert.Len(t, b.failureStreaks, 0)
	b.lock.Unlock()
}

// Test that the benchlist won't bench more than the maximum portion of stake
func TestBenchlistMaxStake(t *testing.T) {
	validators := validation.NewSet()
	validator1 := validation.GenerateRandomValidator(1000)
	validator2 := validation.GenerateRandomValidator(1000)
	validator3 := validation.GenerateRandomValidator(1000)
	validator4 := validation.GenerateRandomValidator(2000)
	validator5 := validation.GenerateRandomValidator(100)
	// Total weight is 5100

	errs := wrappers.Errs{}
	errs.Add(
		validators.AddWeight(validator1.ID(), validator1.Weight()),
		validators.AddWeight(validator2.ID(), validator2.Weight()),
		validators.AddWeight(validator3.ID(), validator3.Weight()),
		validators.AddWeight(validator4.ID(), validator4.Weight()),
		validators.AddWeight(validator5.ID(), validator5.Weight()),
	)
	if errs.Errored() {
		t.Fatal(errs.Err)
	}

	threshold := 3
	duration := 1 * time.Hour
	// Shouldn't bench more than 2550 (5100/2)
	maxPortion := 0.5
	benchIntf, err := NewBenchlist(
		ids.Empty,
		logging.NoLog{},
		&TestBenchable{T: t},
		validators,
		threshold,
		minimumFailingDuration,
		duration,
		maxPortion,
		prometheus.NewRegistry(),
	)
	if err != nil {
		t.Fatal(err)
	}
	b := benchIntf.(*benchlist)
	defer b.timer.Stop()
	now := time.Now()
	b.clock.Set(now)

	// Register [threshold-1] failures for 3 validators
	for _, vdr := range []validation.Validator{validator1, validator2, validator3} {
		for i := 0; i < threshold-1; i++ {
			b.RegisterFailure(vdr.ID())
		}
	}

	// Advance the time to past the minimum failing duration
	newTime := now.Add(minimumFailingDuration).Add(time.Second)
	b.lock.Lock()
	b.clock.Set(newTime)
	b.lock.Unlock()

	// Register another failure for all three
	for _, vdr := range []validation.Validator{validator1, validator2, validator3} {
		b.RegisterFailure(vdr.ID())
	}

	// Only validator1 and validator2 should be benched (total weight 2000)
	// Benching validator3 (weight 1000) would cause the amount benched
	// to exceed the maximum
	b.lock.Lock()
	assert.True(t, b.isBenched(validator1.ID()))
	assert.True(t, b.isBenched(validator2.ID()))
	assert.False(t, b.isBenched(validator3.ID()))
	assert.Equal(t, b.benchedQueue.Len(), 2)
	assert.Equal(t, b.benchlistSet.Len(), 2)
	assert.Len(t, b.failureStreaks, 1)
	fs := b.failureStreaks[validator3.ID()]
	fs.consecutive = threshold
	fs.firstFailure = now
	b.lock.Unlock()

	// Register threshold - 1 failures for validator5
	for i := 0; i < threshold-1; i++ {
		b.RegisterFailure(validator5.ID())
	}

	// Advance the time past min failing duration
	newTime2 := newTime.Add(minimumFailingDuration).Add(time.Second)
	b.lock.Lock()
	b.clock.Set(newTime2)
	b.lock.Unlock()

	// Register another failure for validator5
	b.RegisterFailure(validator5.ID())

	// validator5 should be benched now
	b.lock.Lock()
	assert.True(t, b.isBenched(validator1.ID()))
	assert.True(t, b.isBenched(validator2.ID()))
	assert.True(t, b.isBenched(validator5.ID()))
	assert.Equal(t, 3, b.benchedQueue.Len())
	assert.Equal(t, 3, b.benchlistSet.Len())
	assert.Contains(t, b.benchlistSet, validator1.ID())
	assert.Contains(t, b.benchlistSet, validator2.ID())
	assert.Contains(t, b.benchlistSet, validator5.ID())
	assert.Len(t, b.failureStreaks, 1) // for validator3
	b.lock.Unlock()

	// More failures for validator3 shouldn't add it to the bench
	// because the max bench amount would be exceeded
	for i := 0; i < threshold-1; i++ {
		b.RegisterFailure(validator3.ID())
	}

	b.lock.Lock()
	assert.True(t, b.isBenched(validator1.ID()))
	assert.True(t, b.isBenched(validator2.ID()))
	assert.True(t, b.isBenched(validator5.ID()))
	assert.False(t, b.isBenched(validator3.ID()))
	assert.Equal(t, 3, b.benchedQueue.Len())
	assert.Equal(t, 3, b.benchlistSet.Len())
	assert.Len(t, b.failureStreaks, 1)
	assert.Contains(t, b.failureStreaks, validator3.ID())

	// Ensure the benched queue root has the min end time
	minEndTime := b.benchedQueue[0].benchedUntil
	benchedIDs := []ids.ShortID{validator1.ID(), validator2.ID(), validator5.ID()}
	for _, benchedVdr := range b.benchedQueue {
		assert.Contains(t, benchedIDs, benchedVdr.validatorID)
		assert.True(t, !benchedVdr.benchedUntil.Before(minEndTime))
	}

	b.lock.Unlock()
}

// Test validators are removed from the bench correctly
func TestBenchlistRemove(t *testing.T) {
	validators := validation.NewSet()
	validator1 := validation.GenerateRandomValidator(1000)
	validator2 := validation.GenerateRandomValidator(1000)
	validator3 := validation.GenerateRandomValidator(1000)
	validator4 := validation.GenerateRandomValidator(1000)
	validator5 := validation.GenerateRandomValidator(1000)
	// Total weight is 5100

	errs := wrappers.Errs{}
	errs.Add(
		validators.AddWeight(validator1.ID(), validator1.Weight()),
		validators.AddWeight(validator2.ID(), validator2.Weight()),
		validators.AddWeight(validator3.ID(), validator3.Weight()),
		validators.AddWeight(validator4.ID(), validator4.Weight()),
		validators.AddWeight(validator5.ID(), validator5.Weight()),
	)
	if errs.Errored() {
		t.Fatal(errs.Err)
	}

	count := 0
	benchable := &TestBenchable{
		T:             t,
		CantUnbenched: true,
		UnbenchedF: func(ids.ID, ids.ShortID) {
			count++
		},
	}

	threshold := 3
	duration := 2 * time.Second
	maxPortion := 0.76 // can bench 3 of the 5 validators
	benchIntf, err := NewBenchlist(
		ids.Empty,
		logging.NoLog{},
		benchable,
		validators,
		threshold,
		minimumFailingDuration,
		duration,
		maxPortion,
		prometheus.NewRegistry(),
	)
	if err != nil {
		t.Fatal(err)
	}
	b := benchIntf.(*benchlist)
	defer b.timer.Stop()
	now := time.Now()
	b.lock.Lock()
	b.clock.Set(now)
	b.lock.Unlock()

	// Register [threshold-1] failures for 3 validators
	for _, vdr := range []validation.Validator{validator1, validator2, validator3} {
		for i := 0; i < threshold-1; i++ {
			b.RegisterFailure(vdr.ID())
		}
	}

	// Advance the time past the min failing duration and register another failure
	// for each
	now = now.Add(minimumFailingDuration).Add(time.Second)
	b.lock.Lock()
	b.clock.Set(now)
	b.lock.Unlock()
	for _, vdr := range []validation.Validator{validator1, validator2, validator3} {
		b.RegisterFailure(vdr.ID())
	}

	// All 3 should be benched
	b.lock.Lock()
	assert.True(t, b.isBenched(validator1.ID()))
	assert.True(t, b.isBenched(validator2.ID()))
	assert.True(t, b.isBenched(validator3.ID()))
	assert.Equal(t, 3, b.benchedQueue.Len())
	assert.Equal(t, 3, b.benchlistSet.Len())
	assert.Len(t, b.failureStreaks, 0)

	// Ensure the benched queue root has the min end time
	minEndTime := b.benchedQueue[0].benchedUntil
	benchedIDs := []ids.ShortID{validator1.ID(), validator2.ID(), validator3.ID()}
	for _, benchedVdr := range b.benchedQueue {
		assert.Contains(t, benchedIDs, benchedVdr.validatorID)
		assert.True(t, !benchedVdr.benchedUntil.Before(minEndTime))
	}

	// Set the benchlist's clock past when all validators should be unbenched
	// so that when its timer fires, it can remove them
	b.clock.Set(b.clock.Time().Add(duration))
	b.lock.Unlock()

	// Make sure each validator is eventually removed
	assert.Eventually(
		t,
		func() bool {
			return !b.IsBenched(validator1.ID())
		},
		duration+time.Second, // extra time.Second as grace period
		100*time.Millisecond,
	)

	assert.Eventually(
		t,
		func() bool {
			return !b.IsBenched(validator2.ID())
		},
		duration+time.Second,
		100*time.Millisecond,
	)

	assert.Eventually(
		t,
		func() bool {
			return !b.IsBenched(validator3.ID())
		},
		duration+time.Second,
		100*time.Millisecond,
	)

	assert.Equal(t, 3, count)
}
