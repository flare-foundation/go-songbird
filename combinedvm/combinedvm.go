package combinedvm

import (
	"github.com/flare-foundation/flare/snow/engine/snowman/block"
)

type CombinedVM struct {
	//Vm    *platformvm.VM
	//VmVal *validatorvm.ValidatorVM
	VmVal block.ValidatorVMInterface
	Vm block.ChainVM
}
