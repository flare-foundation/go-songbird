package rpcchainvm

import (
	"context"
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

func (vm *ValidatorsClient) GetValidators(hash []byte) (map[string]float64, error) {
	resp, err := vm.client.GetValidators(context.Background(), &validatorproto.ValidatorsRequest{
		Hash: hash,
	})
	return resp.Validators, err
}

// SetProcess gives ownership of the server process to the client.
func (vm *ValidatorsClient) SetProcess(proc *plugin.Client) {
	vm.proc = proc
}