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

type TestState struct {
	T *testing.T

	CantGetCurrentHeight,
	CantGetValidatorSet bool

	GetCurrentHeightF func() (uint64, error)
	GetValidatorSetF  func(blockID ids.ID) (map[ids.ShortID]uint64, error)
}

func (vm *TestState) GetValidatorSet(blockID ids.ID) (map[ids.ShortID]uint64, error) {
	if vm.GetValidatorSetF != nil {
		return vm.GetValidatorSetF(blockID)
	}
	if vm.CantGetValidatorSet && vm.T != nil {
		vm.T.Fatal(errGetValidatorSet)
	}
	return nil, errGetValidatorSet
}
