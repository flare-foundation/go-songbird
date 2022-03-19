package mocks

import (
	"github.com/flare-foundation/flare/ids"
)

type Updater struct {
	UpdateValidatorsFunc func(blockID ids.ID) error
}

func (m *Updater) UpdateValidators(blockID ids.ID) error {
	return m.UpdateValidatorsFunc(blockID)
}
