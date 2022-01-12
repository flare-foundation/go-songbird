package proposervm

import (
	"errors"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/engine/snowman/block"
)

type ValidatorVM struct {
	*VM
	block.Validators
}

//todo implement where vm calls coreth function to get validators

func (validatorVM *ValidatorVM) GetValidators(hash ids.ID) (map[ids.ID]float64, error) {

	//validatorVM.ChainVM.GetBlock()
	validatorVM.Validators.GetValidators(hash)
	return nil, errors.New("Not implemented fully yet")
}
