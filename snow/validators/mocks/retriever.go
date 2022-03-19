package mocks

import (
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/validators"
)

type Retriever struct {
	GetValidatorsByBlockIDFunc func(blockID ids.ID) (validators.Set, error)
}

func (m *Retriever) GetValidators(blockID ids.ID) (validators.Set, error) {
	return m.GetValidatorsByBlockIDFunc(blockID)
}
