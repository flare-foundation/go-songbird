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

	// GetValidators returns the validator set for the given subnet
	// Returns false if the subnet doesn't exist
	GetValidators() (Set, bool)

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
		validators = coston()
	case constants.SongbirdID:
		validators = songbird()
	case constants.FlareID:
		validators = flare()
	default:
		validators = custom()
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
	validators Set
	maskedVdrs ids.ShortSet
}

// GetValidatorSet implements the Manager interface.
func (m *manager) GetValidators() (Set, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.validators.Len() == 0 {
		return nil, false
	}
	return m.validators, true
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
