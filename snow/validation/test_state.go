// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validation

import (
	"errors"
	"testing"
)

var (
	errCurrentHeight = errors.New("unexpectedly called GetCurrentHeight")
)

type TestState struct {
	T *testing.T

	CantGetCurrentHeight bool

	GetCurrentHeightF func() (uint64, error)
}

func (vm *TestState) GetCurrentHeight() (uint64, error) {
	if vm.GetCurrentHeightF != nil {
		return vm.GetCurrentHeightF()
	}
	if vm.CantGetCurrentHeight && vm.T != nil {
		vm.T.Fatal(errCurrentHeight)
	}
	return 0, errCurrentHeight
}
