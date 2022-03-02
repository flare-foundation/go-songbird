// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcchainvm

import (
	"github.com/hashicorp/go-plugin"
)

type FTSOPlugin struct {
	plugin.NetRPCUnsupportedPlugin
}
