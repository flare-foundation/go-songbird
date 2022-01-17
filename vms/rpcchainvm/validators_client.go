package rpcchainvm

import (
	"context"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow"
	"github.com/flare-foundation/flare/vms/rpcchainvm/validatorproto"
	"github.com/hashicorp/go-plugin"
)

// VMClient is an implementation of VM that talks over RPC.
type ValidatorsClient struct {
	client validatorproto.ValidatorsClient
	broker *plugin.GRPCBroker
	proc   *plugin.Client

	ctx *snow.Context
}

// NewClient returns a VM connected to a remote VM
func NewValidatorsClient(client validatorproto.ValidatorsClient, broker *plugin.GRPCBroker) *ValidatorsClient {
	return &ValidatorsClient{
		client: client,
		broker: broker,
	}
}

func (vm *ValidatorsClient) GetValidators(id ids.ID) (map[ids.ShortID]float64, error) {

	var hash []byte
	copy(hash[:], id.String())
	resp, err := vm.client.GetValidators(context.Background(), &validatorproto.ValidatorsRequest{
		Hash: hash,
	})

	return convertStringMaptoShortIDMap(resp.Validators), err
}

// SetProcess gives ownership of the server process to the client.
func (vm *ValidatorsClient) SetProcess(proc *plugin.Client) {
	vm.proc = proc
}

func convertShortIdmapToStringMap(m map[ids.ShortID]float64) map[string]float64 {
	retM := make(map[string]float64)
	for key, val := range m {
		retM[key.String()] = val
	}
	return retM
}

func convertStringMaptoShortIDMap(m map[string]float64) map[ids.ShortID]float64 {
	retM := make(map[ids.ShortID]float64)
	for key, val := range m {
		retM[stringToShortID(key)] = val
	}
	return retM
}

func stringToShortID(s string) ids.ShortID {
	var shortId [20]byte
	copy(shortId[:], s)
	return shortId
}
