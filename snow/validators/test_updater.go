// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"github.com/flare-foundation/flare/ids"
)

type TestUpdater struct {
	UpdateValidatorsFunc func(blockID ids.ID) error
}

func (m *TestUpdater) UpdateValidators(blockID ids.ID) error {
	return m.UpdateValidatorsFunc(blockID)
}
