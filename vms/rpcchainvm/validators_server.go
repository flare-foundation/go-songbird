// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcchainvm

import (
	"context"
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
	valVM  block.Validators
	broker *plugin.GRPCBroker

	serverCloser grpcutils.ServerCloser
	connCloser   wrappers.Closer

	ctx    *snow.Context
	closed chan struct{}
}

// ValidatorsServer returns a vm instance connected to a remote vm instance
func NewValidatorsServer(valVM block.Validators, broker *plugin.GRPCBroker) *ValidatorsServer {
	return &ValidatorsServer{
		valVM:  valVM,
		broker: broker,
	}
}

func (v ValidatorsServer) GetValidators(_ context.Context, req *validatorproto.ValidatorsRequest) (*validatorproto.ValidatorsResponse, error) {
	validators, err := v.valVM.GetValidators(req.Hash)
	return &validatorproto.ValidatorsResponse{
		Validators: validators, //todo get the map types consistent, maybe just string for now?
	}, err
}
