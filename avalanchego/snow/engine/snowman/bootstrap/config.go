// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrap

import (
	"github.com/flare-foundation/flare/snow/engine/common"
	"github.com/flare-foundation/flare/snow/engine/common/queue"
	"github.com/flare-foundation/flare/snow/engine/common/tracker"
	"github.com/flare-foundation/flare/snow/engine/snowman/block"
)

type Config struct {
	common.Config
	common.AllGetsServer

	// Blocked tracks operations that are blocked on blocks
	Blocked *queue.JobsWithMissing

	VM            block.ChainVM
	WeightTracker tracker.WeightTracker

	Bootstrapped func()
}
