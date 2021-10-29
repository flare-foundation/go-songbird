package nftfx

import (
	"github.com/flare-foundation/flare/vms/secp256k1fx"
)

type Credential struct {
	secp256k1fx.Credential `serialize:"true"`
}
