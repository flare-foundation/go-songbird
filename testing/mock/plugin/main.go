// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"github.com/hashicorp/go-plugin"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/validation"
	"github.com/flare-foundation/flare/testing/mock/chainvm"
	"github.com/flare-foundation/flare/vms/rpcchainvm"
)

func main() {
	mockSet := validation.NewSet()
	for i := 1; i <= 10; i++ {
		_ = mockSet.AddWeight(fakeShortID(i), uint64(i))
	}

	mock := chainvm.ChainVMMock{
		GetValidatorsFunc: func(_ ids.ID) (validation.Set, error) {
			return mockSet, nil
		},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: rpcchainvm.Handshake,
		Plugins: map[string]plugin.Plugin{
			"vm": rpcchainvm.New(&mock),
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}

func fakeShortID(i int) ids.ShortID {
	return ids.ShortID{byte(i), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}
