package propertyfx

import (
	"github.com/flare-foundation/flare/vms/secp256k1fx"
)

type OwnedOutput struct {
	secp256k1fx.OutputOwners `serialize:"true"`
}
