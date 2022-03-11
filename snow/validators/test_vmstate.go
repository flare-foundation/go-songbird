package validators

import (
	"errors"
	"testing"
)

var (
	errGetCurrentHeight = errors.New("unexpectedly called GetCurrentHeight")
)

type TestVMState struct {
	T *testing.T

	GetCurrentHeightF func() (uint64, error)
}

func (T *TestVMState) GetCurrentHeight() (uint64, error) {
	if T.GetCurrentHeightF != nil {
		return T.GetCurrentHeightF()
	}

	return 0, errGetCurrentHeight
}
