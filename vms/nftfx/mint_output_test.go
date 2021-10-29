package nftfx

import (
	"testing"

	"github.com/flare-foundation/flare/vms/components/verify"
)

func TestMintOutputState(t *testing.T) {
	intf := interface{}(&MintOutput{})
	if _, ok := intf.(verify.State); !ok {
		t.Fatalf("should be marked as state")
	}
}
