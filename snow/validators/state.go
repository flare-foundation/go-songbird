// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"sync"

	"github.com/flare-foundation/flare/ids"
)

var _ State = &lockedState{}

// State allows the lookup of the validator set at the requested block ID.
type State interface {

	// GetValidatorSet returns the weights of the nodeIDs at the requested block
	// ID. The returned map should not be modified.
	GetValidatorSet(blockID ids.ID) (map[ids.ShortID]uint64, error)
}

type lockedState struct {
	lock sync.Locker
	s    State
}

func NewLockedState(lock sync.Locker, s State) State {
	return &lockedState{
		lock: lock,
		s:    s,
	}
}

func (s *lockedState) GetValidatorSet(blockID ids.ID) (map[ids.ShortID]uint64, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.s.GetValidatorSet(blockID)
}

type noState struct{}

func NewNoState() State {
	return &noState{}
}

func (s *noState) GetCurrentHeight() (uint64, error) {
	return 0, nil
}

func (s *noState) GetValidatorSet(blockID ids.ID) (map[ids.ShortID]uint64, error) {
	return nil, nil
}
