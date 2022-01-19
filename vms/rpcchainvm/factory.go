// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcchainvm

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/flare-foundation/flare/snow"
	"github.com/flare-foundation/flare/utils/subprocess"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

var errWrongVM = errors.New("wrong vm type")

type Factory struct {
	Path string
}

func (f *Factory) New(ctx *snow.Context) (interface{}, error) {
	// Ignore warning from launching an executable with a variable command
	// because the command is a controlled and required input

	fmt.Println("Factory New called", f.Path)
	config := &plugin.ClientConfig{
		HandshakeConfig: Handshake,
		Plugins:         PluginMap,
		Cmd:             subprocess.New(f.Path),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC,
			plugin.ProtocolGRPC,
		},
		// We kill this client by calling kill() when the chain running this VM
		// shuts down. However, there are some cases where the VM's Shutdown
		// method is not called. Namely, if:
		// 1) The node shuts down after the client is created but before the
		//    chain is registered with the message router.
		// 2) The chain doesn't handle a shutdown message before the node times
		//    out on the chain's shutdown and dies, leaving the shutdown message
		//    unhandled.
		// We set managed to true so that we can call plugin.CleanupClients on
		// node shutdown to ensure every plugin subprocess is killed.
		Managed: true,
	}
	if ctx != nil {
		log.SetOutput(ctx.Log)
		config.Stderr = ctx.Log
		config.Logger = hclog.New(&hclog.LoggerOptions{
			Output: ctx.Log,
			Level:  hclog.Info,
		})
	} else {
		log.SetOutput(ioutil.Discard)
		config.Stderr = ioutil.Discard
		config.Logger = hclog.New(&hclog.LoggerOptions{
			Output: ioutil.Discard,
		})
	}
	client := plugin.NewClient(config)
	fmt.Println("client is: ", client)

	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("Factory New error 1")
		client.Kill()
		return nil, err
	}

	raw, err := rpcClient.Dispense("vm")
	if err != nil {
		fmt.Println("Factory New error 2")
		client.Kill()
		return nil, err
	}

	vm, ok := raw.(*VMClient)
	if !ok {
		client.Kill()
		return nil, errWrongVM
	}

	vm.SetProcess(client) // todo do the same for validators

	vm.ctx = ctx
	//
	raw1, err := rpcClient.Dispense("validators")
	if err != nil {
		fmt.Println("Factory New error 3")
		client.Kill()
		return nil, err
	}
	valVM, ok := raw1.(*ValidatorsClient)
	if !ok {
		fmt.Println("Factory New error 4")
		client.Kill()
		return nil, errWrongVM
	}
	valVM.SetProcess(client)
	valVM.ctx = ctx
	GlobalValidatorClient = valVM
	//return []interface{}{vm, valVM}, nil
	return vm, nil
}
