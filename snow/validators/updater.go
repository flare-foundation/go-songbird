// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"fmt"

	"github.com/flare-foundation/flare/ids"
)

type Updater interface {
	UpdateValidators(blockID ids.ID) error
}

func NewUpdater(validators Set, retriever Retriever) Updater {
	u := updater{
		validators: validators,
		retriever:  retriever,
	}
	return &u
}

type updater struct {
	validators Set
	retriever  Retriever
}

func (u *updater) UpdateValidators(blockID ids.ID) error {
	validators, err := u.retriever.GetValidators(blockID)
	if err != nil {
		return fmt.Errorf("could not get validators (block: %x): %w", blockID, err)
	}
	err = u.validators.Set(validators.List())
	if err != nil {
		return fmt.Errorf("could not set validator set: %w", err)
	}
	return nil
}
