// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/consensus/snowman"
	"github.com/flare-foundation/flare/snow/engine/common"
	"github.com/flare-foundation/flare/utils/wrappers"
)

// convincer sends chits to [vdr] once all its dependencies are met
type convincer struct {
	consensus snowman.Consensus
	sender    common.Sender
	vdr       ids.ShortID
	requestID uint32
	sent      bool
	abandoned bool
	deps      ids.Set
	errs      *wrappers.Errs
}

func (c *convincer) Dependencies() ids.Set { return c.deps }

// Mark that a dependency has been met
func (c *convincer) Fulfill(id ids.ID) {
	c.deps.Remove(id)
	c.Update()
}

// Abandon this attempt to send chits.
func (c *convincer) Abandon(ids.ID) {
	c.abandoned = true
	c.Update()
}

func (c *convincer) Update() {
	if c.sent || c.errs.Errored() || (!c.abandoned && c.deps.Len() != 0) {
		return
	}
	c.sent = true

	pref := []ids.ID{c.consensus.Preference()}
	c.sender.SendChits(c.vdr, c.requestID, pref)
}
