package propertyfx

import (
	"github.com/flare-foundation/flare/snow"
	"github.com/flare-foundation/flare/vms/components/verify"
	"github.com/flare-foundation/flare/vms/secp256k1fx"
)

type BurnOperation struct {
	secp256k1fx.Input `serialize:"true"`
}

func (op *BurnOperation) InitCtx(ctx *snow.Context) {}

func (op *BurnOperation) Outs() []verify.State { return nil }
