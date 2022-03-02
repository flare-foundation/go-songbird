// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"errors"
	"testing"

	"github.com/flare-foundation/flare/ids"
)

var (
	errGetValidatorSet = errors.New("unexpectedly called GetValidatorSet")
)

type TestManager struct {
	T *testing.T

	CantGetValidatorSet          bool
	CantGetValidatorSetByBlockID bool

	GetValidatorSetF          func() (map[ids.ShortID]uint64, error)
	GetValidatorSetByBlockIDF func(blockID ids.ID) (map[ids.ShortID]uint64, error)
}

func (tm *TestManager) GetValidatorSet(blockID ids.ID) (map[ids.ShortID]uint64, error) {
	if tm.GetValidatorSetF != nil {
		return tm.GetValidatorSetF()
	}
	if tm.CantGetValidatorSet && tm.T != nil {
		tm.T.Fatal(errGetValidatorSet)
	}
	return nil, errGetValidatorSet
}

func (tm *TestManager) GetValidatorSetByBlockID(blockID ids.ID) (map[ids.ShortID]uint64, error) {
	if tm.GetValidatorSetF != nil {
		return tm.GetValidatorSetByBlockIDF(blockID)
	}
	if tm.CantGetValidatorSet && tm.T != nil {
		tm.T.Fatal(errGetValidatorSet)
	}
	return nil, errGetValidatorSet
}
