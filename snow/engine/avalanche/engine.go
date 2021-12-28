// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avalanche

import (
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/consensus/avalanche"
	"github.com/flare-foundation/flare/snow/engine/common"
)

// Engine describes the events that can occur on a consensus instance
type Engine interface {
	common.Engine

	// Initialize this engine.
	Initialize(Config) error

	// GetVtx returns a vertex by its ID.
	// Returns an error if unknown.
	GetVtx(vtxID ids.ID) (avalanche.Vertex, error)
}
