package validatorvm

import (
	"fmt"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/engine/snowman/block"
)

type ValidatorVM struct {
	//VM proposervm.VM
	block.ValidatorVMInterface // todo how to initialize it?
	//platformvm.Factory
}

//todo implement where vm calls coreth function to get validators

func New(valVM block.ValidatorVMInterface) *ValidatorVM {
	return &ValidatorVM{
		valVM,
	}
}

func (validatorVM *ValidatorVM) GetValidators(hash ids.ID) (map[ids.ShortID]float64, error) { // todo it takes ids.ID but returns map where key is shortID
	fmt.Println("test..")
	//validatorVM.ChainVM.GetBlock()
	fmt.Println("Calling GetValidators() in validatorvm/validatorvm.go")

	valMap, err := validatorVM.ValidatorVMInterface.GetValidators(hash) // todo this is where it calls the underlying coreth call?
	return valMap, err
	//return nil, errors.New("Not implemented fully yet")
}
