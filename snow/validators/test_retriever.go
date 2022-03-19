// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"github.com/flare-foundation/flare/ids"
)

type TestRetriever struct {
	GetValidatorsByBlockIDFunc func(blockID ids.ID) (Set, error)
}

func (m *TestRetriever) GetValidators(blockID ids.ID) (Set, error) {
	return m.GetValidatorsByBlockIDFunc(blockID)
}
