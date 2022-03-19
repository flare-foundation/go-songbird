package validators

import (
	"github.com/flare-foundation/flare/ids"
)

type UpdaterMock struct {
	UpdateValidatorsF func(blockID ids.ID) error
}

func (m *UpdaterMock) UpdateValidators(blockID ids.ID) error {
	return m.UpdateValidatorsF(blockID)

}
