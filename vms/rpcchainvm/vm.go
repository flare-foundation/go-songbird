// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcchainvm

import (
	"github.com/flare-foundation/flare/vms/rpcchainvm/validatorproto"
	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/hashicorp/go-plugin"

	"github.com/flare-foundation/flare/snow/engine/snowman/block"
	"github.com/flare-foundation/flare/vms/rpcchainvm/vmproto"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  8,
	MagicCookieKey:   "VM_PLUGIN",
	MagicCookieValue: "dynamic",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"vm":         &Plugin{},
	"validators": &PluginValidator{},
}

var GlobalValidatorClient *ValidatorsClient
var GlobalVMClient *VMClient

// Plugin is the implementation of plugin.Plugin so we can serve/consume this.
// We also implement GRPCPlugin so that this plugin can be served over gRPC.
type Plugin struct {
	plugin.NetRPCUnsupportedPlugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	vm block.ChainVM
}

// New creates a new plugin from the provided VM
func New(vm block.ChainVM) *Plugin { return &Plugin{vm: vm} }

// GRPCServer registers a new GRPC server.
func (p *Plugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	vmproto.RegisterVMServer(s, NewServer(p.vm, broker))
	return nil
}

// GRPCClient returns a new GRPC client
func (p *Plugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return NewClient(vmproto.NewVMClient(c), broker), nil
}

type PluginValidator struct {
	//plugin.NetRPCUnsupportedPlugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	// GRPCPlugin must still implement the Plugin interface
    plugin.Plugin
	ValVM     block.ValidatorVMInterface // proposervm.ValidatorVM
}

// NewPluginValidator creates a new PluginValidator from the provided VM
func NewPluginValidator(vm block.ValidatorVMInterface) *PluginValidator {
	return &PluginValidator{ValVM: vm}
}

// GRPCServer registers a new GRPC server.
func (p *PluginValidator) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	validatorproto.RegisterValidatorsServer(s, NewValidatorsServer(p.ValVM, broker)) //todo
	return nil
}

// GRPCClient returns a new GRPC client
func (p *PluginValidator) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return NewValidatorsClient(validatorproto.NewValidatorsClient(c), broker), nil
}
