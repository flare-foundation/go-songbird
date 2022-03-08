package validators

import (
	"fmt"

	"github.com/flare-foundation/flare/ids"
)

type Updater interface {
	UpdateValidators(blockID ids.ID) error
}

func NewUpdater(validators Set, retriever Retriever) Updater {
	fmt.Println("NewUpdater called.")
	validators.AddWeight(ids.ShortID{1}, 123)
	u := updater{
		validators: validators,
		retriever:  retriever,
	}
	fmt.Println(u)
	return &u
}

type updater struct {
	validators Set
	retriever  Retriever
}

func (u *updater) UpdateValidators(blockID ids.ID) error {
	validators, err := u.retriever.GetValidatorsByBlockID(blockID)
	if err != nil {
		return fmt.Errorf("could not get validators (block: %x): %w", blockID, err)
	}
	err = u.validators.Set(validators.List())
	if err != nil {
		return fmt.Errorf("could not set validator set: %w", err)
	}
	return nil
}
