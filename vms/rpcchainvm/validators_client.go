package rpcchainvm

//import (
//
//	"github.com/flare-foundation/flare/vms/rpcchainvm/vmproto"
//	"github.com/hashicorp/go-plugin"
//)
//
//// VMClient is an implementation of VM that talks over RPC.
//type ValidatorsClient struct {
//
//}
//
//// NewClient returns a VM connected to a remote VM
//func NewClient(client vmproto.VMClient, broker *plugin.GRPCBroker) *VMClient {
//	return &VMClient{
//		client: client,
//		broker: broker,
//	}
//}
//
//// SetProcess gives ownership of the server process to the client.
//func (vm *VMClient) SetProcess(proc *plugin.Client) {
//	vm.proc = proc
//}
