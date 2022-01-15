package proposervm

import (
	"errors"
	"fmt"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/engine/snowman/block"
)

type ValidatorVM struct {
	*VM
	block.ValidatorVMInterface
}

//todo implement where vm calls coreth function to get validators

func (validatorVM *ValidatorVM) GetValidators(hash ids.ID) (map[ids.ID]float64, error) {

	//validatorVM.ChainVM.GetBlock()
	fmt.Println("Calling GetValidators() in proposervm/validatorvm")

	validatorVM.ValidatorVMInterface.GetValidators(hash)
	return nil, errors.New("Not implemented fully yet")
}
