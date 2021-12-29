// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package version

import (
	"errors"
	"time"

	"github.com/flare-foundation/flare/utils/timer/mockable"
)

var (
	errIncompatible               = errors.New("peers version is incompatible")
	errMaskable                   = errors.New("peers version is maskable")
	_               Compatibility = &compatibility{}
)

// Compatibility a utility for checking the compatibility of peer versions
type Compatibility interface {
	// Returns the local version
	Version() Application

	// Returns nil if the provided version is compatible with the local version.
	// This means that the version is connectable and that consensus messages
	// can be made to them.
	Compatible(Application) error

	// Returns nil if the provided version shouldn't be masked. This means that
	// the version is connectable but not compatible. The version is so old that
	// it should just be masked.
	Unmaskable(Application) error

	// Returns nil if the provided version will not be masked by this version.
	WontMask(Application) error

	// Returns when additional masking will occur.
	MaskTime() time.Time
}

type compatibility struct {
	version Application

	minCompatible     Application
	minCompatibleTime time.Time
	prevMinCompatible Application

	minUnmaskable     Application
	minUnmaskableTime time.Time
	prevMinUnmaskable Application

	clock mockable.Clock
}

// NewCompatibility returns a compatibility checker with the provided options
func NewCompatibility(
	version Application,
	minCompatible Application,
	minCompatibleTime time.Time,
	prevMinCompatible Application,
	minUnmaskable Application,
	minUnmaskableTime time.Time,
	prevMinUnmaskable Application,
) Compatibility {
	return &compatibility{
		version:           version,
		minCompatible:     minCompatible,
		minCompatibleTime: minCompatibleTime,
		prevMinCompatible: prevMinCompatible,
		minUnmaskable:     minUnmaskable,
		minUnmaskableTime: minUnmaskableTime,
		prevMinUnmaskable: prevMinUnmaskable,
	}
}

func (c *compatibility) Version() Application { return c.version }

func (c *compatibility) Compatible(peer Application) error {
	// NOTE: This hard-coded Flare versioning support ensures that we don't drop
	// newer nodes which send Flare application name and version numbers.
	// TODO: Remove when all nodes on Songbird are sending Flare application
	// name and version numbers.
	if peer.App() == "flare" {
		return nil
	}

	if err := c.version.Compatible(peer); err != nil {
		return err
	}

	if !peer.Before(c.minCompatible) {
		// The peer is at least the minimum compatible version.
		return nil
	}

	// The peer is going to be marked as incompatible at [c.minCompatibleTime].
	now := c.clock.Time()
	if !now.Before(c.minCompatibleTime) {
		return errIncompatible
	}

	// The minCompatible check isn't being enforced yet.
	if !peer.Before(c.prevMinCompatible) {
		// The peer is at least the previous minimum compatible version.
		return nil
	}
	return errIncompatible
}

func (c *compatibility) Unmaskable(peer Application) error {
	// NOTE: This hard-coded Flare versioning support ensures that we don't drop
	// newer nodes which send Flare application name and version numbers.
	// TODO: Remove when all nodes on Songbird are sending Flare application
	// name and version numbers.
	if peer.App() == "flare" {
		return nil
	}

	if err := c.Compatible(peer); err != nil {
		return err
	}

	if !peer.Before(c.minUnmaskable) {
		// The peer is at least the minimum unmaskable version.
		return nil
	}

	// The peer is going to be marked as maskable at [c.minUnmaskableTime].
	now := c.clock.Time()
	if !now.Before(c.minUnmaskableTime) {
		return errMaskable
	}

	// The minCompatible check isn't being enforced yet.
	if !peer.Before(c.prevMinUnmaskable) {
		// The peer is at least the previous minimum unmaskable version.
		return nil
	}
	return errMaskable
}

func (c *compatibility) WontMask(peer Application) error {
	// NOTE: This hard-coded Flare versioning support ensures that we don't drop
	// newer nodes which send Flare application name and version numbers.
	// TODO: Remove when all nodes on Songbird are sending Flare application
	// name and version numbers.
	if peer.App() == "flare" {
		return nil
	}

	if err := c.Compatible(peer); err != nil {
		return err
	}

	if !peer.Before(c.minUnmaskable) {
		// The peer is at least the minimum unmaskable version.
		return nil
	}
	return errMaskable
}

func (c *compatibility) MaskTime() time.Time {
	return c.minUnmaskableTime
}
