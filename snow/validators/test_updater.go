package validators

import (
	"errors"
	"testing"

	"github.com/flare-foundation/flare/ids"
)

var (
	errUpdateValidators = errors.New("unexpectedly called UpdateValidators")
)

type TestUpdater struct {
	T *testing.T

	UpdateValidatorsF func(blockID ids.ID) error
}

func (T *TestUpdater) UpdateValidators(blockID ids.ID) error {
	if T.UpdateValidatorsF != nil {
		return T.UpdateValidatorsF(blockID)
	}

	return errUpdateValidators
}
