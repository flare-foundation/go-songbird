// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validation

import (
	"fmt"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"
)

type Updater interface {
	UpdateValidators(blockID ids.ID) error
}

func NewRetrievingUpdater(log logging.Logger, validators Set, retrieve Retriever) *RetrievingUpdater {
	u := RetrievingUpdater{
		log:        log,
		validators: validators,
		retrieve:   retrieve,
	}
	return &u
}

type RetrievingUpdater struct {
	log        logging.Logger
	validators Set
	retrieve   Retriever
}

func (u *RetrievingUpdater) UpdateValidators(blockID ids.ID) error {
	validators, err := u.retrieve.GetValidators(blockID)
	if err != nil {
		return fmt.Errorf("could not get validators for updating: %w", err)
	}
	err = u.validators.Set(validators.List())
	if err != nil {
		return fmt.Errorf("could not set validators: %w", err)
	}
	u.log.Debug("validators updated: %s", validators)
	return nil
}
