// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package poll

import (
	"fmt"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/formatting"
)

// Set is a collection of polls
type Set interface {
	fmt.Stringer

	Add(requestID uint32, validators ids.ShortBag) bool
	Vote(requestID uint32, validator ids.ShortID, vote ids.ID) []ids.Bag
	Drop(requestID uint32, validator ids.ShortID) []ids.Bag
	Len() int
}

// Poll is an outstanding poll
type Poll interface {
	formatting.PrefixedStringer

	Vote(validator ids.ShortID, vote ids.ID)
	Drop(validator ids.ShortID)
	Finished() bool
	Result() ids.Bag
}

// Factory creates a new Poll
type Factory interface {
	New(validators ids.ShortBag) Poll
}
