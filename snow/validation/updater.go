// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validation

import (
	"fmt"

	"github.com/flare-foundation/flare/ids"
)

type Updater interface {
	UpdateValidators(blockID ids.ID) error
}

func NewRetrievingUpdater(validators Set, retriever Retriever) *RetrievingUpdater {
	u := RetrievingUpdater{
		validators: validators,
		retriever:  retriever,
	}
	return &u
}

type RetrievingUpdater struct {
	validators Set
	retriever  Retriever
}

func (u *RetrievingUpdater) UpdateValidators(blockID ids.ID) error {
	validators, err := u.retriever.GetValidators(blockID)
	if err != nil {
		return fmt.Errorf("could not get validators for updating: %w", err)
	}
	err = u.validators.Set(validators.List())
	if err != nil {
		return fmt.Errorf("could not set validators: %w", err)
	}
	return nil
}
