package rpcchainvm

import (
	"context"
	"fmt"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/vms/rpcchainvm/validatorproto"
	"github.com/hashicorp/go-plugin"

	"github.com/flare-foundation/flare/snow"
	"github.com/flare-foundation/flare/snow/engine/snowman/block"
	"github.com/flare-foundation/flare/utils/wrappers"
	"github.com/flare-foundation/flare/vms/rpcchainvm/grpcutils"
)

var _ validatorproto.ValidatorsServer = &ValidatorsServer{}

// ValidatorsServer is a VM that is managed over RPC.
type ValidatorsServer struct {
	validatorproto.UnimplementedValidatorsServer
	valVM  block.ValidatorVMInterface
	broker *plugin.GRPCBroker

	serverCloser grpcutils.ServerCloser
	connCloser   wrappers.Closer

	ctx    *snow.Context
	closed chan struct{}
}

// ValidatorsServer returns a vm instance connected to a remote vm instance
func NewValidatorsServer(valVM block.ValidatorVMInterface, broker *plugin.GRPCBroker) *ValidatorsServer {
	return &ValidatorsServer{
		valVM:  valVM,
		broker: broker,
	}
}

func (v ValidatorsServer) GetValidators(_ context.Context, req *validatorproto.ValidatorsRequest) (*validatorproto.ValidatorsResponse, error) {
	bytesID, err := ids.ToID(req.Hash)
	if err != nil {
		return nil, err
	}
	fmt.Println("Calling GetValidators() in validators_server.go")

	validators, err := v.valVM.GetValidators(bytesID)
	return &validatorproto.ValidatorsResponse{
		Validators: convertMapIDstoMapStringKey(validators), //todo get the map types consistent, maybe just string for now?
	}, err
}

func convertMapIDstoMapStringKey(m map[ids.ShortID]float64) map[string]float64 {
	stringMap := make(map[string]float64)
	for key, value := range m {
		stringMap[key.String()] = value
	}
	return stringMap
}