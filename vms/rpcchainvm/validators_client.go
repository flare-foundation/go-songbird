package rpcchainvm

import (
	"context"
	"fmt"
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
	fmt.Println("Printing hash and id")

	fmt.Println(id.String(), hash)
	hash = []byte(id.String())
	fmt.Println(hash)
	fmt.Println("id[:]: ", id[:])
	var hash2 [32]byte
	hash3 := hash2[:]
	resp, err := vm.client.GetValidators(context.Background(), &validatorproto.ValidatorsRequest{
		Hash: hash3,
	})
	if err != nil {
		fmt.Println("Error in GetValidators of ValidatorsClient: ", err.Error())
	}
	fmt.Println("Response in GetValidators of ValidatorsClient: ", resp)
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
