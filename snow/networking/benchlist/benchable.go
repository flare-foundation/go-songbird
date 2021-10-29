// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package benchlist

import (
	"github.com/flare-foundation/flare/ids"
)

// Benchable is notified when a validator is benched or unbenched from a given chain
type Benchable interface {
	// Mark that [validatorID] has been benched on the given chain
	Benched(chainID ids.ID, validatorID ids.ShortID)
	// Mark that [validatorID] has been unbenched from the given chain
	Unbenched(chainID ids.ID, validatorID ids.ShortID)
}
