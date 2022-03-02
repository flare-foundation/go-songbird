// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"fmt"
	"strings"
	"sync"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
)

// Manager holds the validator set of each subnet
type Manager interface {
	fmt.Stringer

	// SetSource sets the dynamic source of validators for this manager.
	SetSource(source Source) error

	// GetValidators returns the latest validator set.
	GetValidators() (Set, error)

	// GetValidatorsByBlockID returns the validator set
	GetValidatorsByBlockID(blockID ids.ID) (Set, error)

	// MaskValidator hides the named validator from future samplings
	MaskValidator(vdrID ids.ShortID) error

	// RevealValidator ensures the named validator is not hidden from future
	// samplings
	RevealValidator(vdrID ids.ShortID) error

	// Contains returns true if there is a validator with the specified ID
	// currently in the set.
	Contains(vdrID ids.ShortID) bool
}

// NewManager returns a new, empty manager
func NewManager(networkID uint32) Manager {
	var validators Set
	switch networkID {
	case constants.CostonID:
		validators = loadCostonValidators()
	case constants.SongbirdID:
		validators = loadSongbirdValidators()
	case constants.FlareID:
		validators = loadFlareValidators()
	default:
		validators = loadCustomValidators()
	}
	return &manager{
		networkID:  networkID,
		validators: validators,
	}
}

// manager implements Manager
type manager struct {
	lock       sync.Mutex
	networkID  uint32
	lastID     ids.ID
	validators Set
	cache      map[ids.ID]Set
	maskedVdrs ids.ShortSet
	source     Source
}

// SetSource sets the source that we load the dynamic set of validators from.
func (m *manager) SetSource(source Source) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.source != nil {
		return fmt.Errorf("validator source already set")
	}
	m.source = source
	return nil
}

// GetValidators implements the validator manager interface. It should _always_
// return the same set (same memory location), because it is used to propagate
// the recent set of validators across many components.
func (m *manager) GetValidators() (Set, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	// If source has not been set, validators can be empty if people forgot to
	// set `CUSTOM_VALIDATORS` on a custom network. Otherwise it will contain
	// the default set for the respective network.
	if m.source == nil && m.validators.Len() == 0 {
		return nil, ErrNoValidators
	}

	// If source is not set, but we have a default set of validators with the
	// network, we can return the default set.
	if m.source == nil {
		return m.validators, nil
	}

	// If the source is set, we can get the accepted block ID. If the accepted
	// ID has not changed, we can simply return the current validator set.
	acceptedID, err := m.source.LastAccepted()
	if err != nil {
		return nil, fmt.Errorf("could not get accepted block: %w", err)
	}
	if acceptedID == m.lastID {
		return m.validators, nil
	}

	// At this point, the last ID has changed.
	m.lastID = acceptedID

	// Otherwise, we check if this validator set has already been loaded from
	// the EVM before through a specific retrieval by block ID. In that case, we
	// update the latest validator set and return.
	set, ok := m.cache[acceptedID]
	if ok {
		err = m.validators.Set(set.List())
		if err != nil {
			return nil, fmt.Errorf("could not set latest validator set: %w", err)
		}
		return m.validators, nil
	}

	// Last but not least, if this is the first attempt to get this validator
	// set, we retrieve it from the EVM, cache it under the accepted ID and
	// update the latest set to have the same content.
	set, err = m.getValidatorsByBlockID(acceptedID)
	if err != nil {
		return nil, fmt.Errorf("could not get validators (accepted: %x): %w", acceptedID, err)
	}
	m.cache[acceptedID] = set

	err = m.validators.Set(set.List())
	if err != nil {
		return nil, fmt.Errorf("could not set latest validator set: %w", err)
	}

	return m.validators, nil
}

// GetValidatorsByBlockID implements the Manager interface.
func (m *manager) GetValidatorsByBlockID(blockID ids.ID) (Set, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	// If no source is set, we should error.
	if m.source == nil {
		return nil, fmt.Errorf("validator source not set yet")
	}

	// If this set is already cached, we return it.
	set, ok := m.cache[blockID]
	if ok {
		return set, nil
	}

	// Otherwise, we try to load it from the EVM.
	set, err := m.getValidatorsByBlockID(blockID)
	if err != nil {
		return nil, fmt.Errorf("could not get validators (block: %x): %w", blockID, err)
	}

	// Cache and return after loading form EVM.
	m.cache[blockID] = set

	return set, nil
}

// getValidatorsByBlockID uses the injected
func (m *manager) getValidatorsByBlockID(blockID ids.ID) (Set, error) {

	// This call loads the validators map from the EVM RPC client.
	validatorMap, err := m.source.LoadValidators(blockID)
	if err != nil {
		return nil, fmt.Errorf("could not load validators: %w", err)
	}

	// Convert into validator set and return.
	set := NewSet()
	for validatorID, weight := range validatorMap {
		err = set.AddWeight(validatorID, weight)
		if err != nil {
			return nil, fmt.Errorf("could not set validator weight (validator: %x, weight: %d): %w", validatorID, weight, err)
		}
	}

	return set, nil
}

// MaskValidator implements the Manager interface.
func (m *manager) MaskValidator(vdrID ids.ShortID) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.maskedVdrs.Contains(vdrID) {
		return nil
	}
	m.maskedVdrs.Add(vdrID)

	if err := m.validators.MaskValidator(vdrID); err != nil {
		return err
	}
	return nil
}

// RevealValidator implements the Manager interface.
func (m *manager) RevealValidator(vdrID ids.ShortID) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if !m.maskedVdrs.Contains(vdrID) {
		return nil
	}
	m.maskedVdrs.Remove(vdrID)

	if err := m.validators.RevealValidator(vdrID); err != nil {
		return err
	}
	return nil
}

// Contains implements the Manager interface.
func (m *manager) Contains(vdrID ids.ShortID) bool {
	m.lock.Lock()
	defer m.lock.Unlock()

	return m.validators.Contains(vdrID)
}

func (m *manager) String() string {
	m.lock.Lock()
	defer m.lock.Unlock()
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("Validator Set: (Size = %d)",
		m.validators.Len(),
	))
	sb.WriteString(fmt.Sprintf(
		"\n    Network[%d]: %s",
		m.networkID,
		m.validators.PrefixedString("    "),
	))

	return sb.String()
}
