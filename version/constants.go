// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package version

import (
	"time"

	"github.com/flare-foundation/flare/utils/constants"
)

// NOTE: We are doing a three-phase deploy to get rid of legacy Avalanche
// versioning and to transition to Flare versioning.
// 1. Roll out upgrade with hard-coded Flare versioning support.
// 2. Roll out upgrade sending Flare application name and version numbers.
// 3. Roll out upgrade removing hard-coded Flare versioning support.

// These are globals that describe network upgrades and node versions
var (
	// TODO: Change to Flare versions after all nodes on Songbird have hard-coded
	// Flare versioning support.
	Current                      = NewDefaultVersion(1, 7, 2)
	CurrentApp                   = NewDefaultApplication(constants.PlatformName, Current.Major(), Current.Minor(), Current.Patch())
	MinimumCompatibleVersion     = NewDefaultApplication(constants.PlatformName, 1, 7, 0)
	PrevMinimumCompatibleVersion = NewDefaultApplication(constants.PlatformName, 1, 6, 0)
	MinimumUnmaskedVersion       = NewDefaultApplication(constants.PlatformName, 1, 1, 0)
	PrevMinimumUnmaskedVersion   = NewDefaultApplication(constants.PlatformName, 1, 0, 0)
	VersionParser                = NewDefaultApplicationParser()

	CurrentDatabase = DatabaseVersion1_4_5
	PrevDatabase    = DatabaseVersion1_0_0

	DatabaseVersion1_4_5 = NewDefaultVersion(1, 4, 5)
	DatabaseVersion1_0_0 = NewDefaultVersion(1, 0, 0)

	ApricotPhase0Times       = map[uint32]time.Time{}
	ApricotPhase0DefaultTime = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)

	ApricotPhase1Times       = map[uint32]time.Time{}
	ApricotPhase1DefaultTime = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)

	ApricotPhase2Times       = map[uint32]time.Time{}
	ApricotPhase2DefaultTime = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)

	ApricotPhase3Times = map[uint32]time.Time{
		constants.CostonID:   time.Date(2100, time.January, 1, 0, 0, 0, 0, time.UTC),
		constants.SongbirdID: time.Date(2100, time.January, 1, 0, 0, 0, 0, time.UTC),
		constants.FlareID:    time.Date(2100, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	ApricotPhase3DefaultTime = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)

	ApricotPhase4Times = map[uint32]time.Time{
		constants.CostonID:   time.Date(2100, time.January, 1, 0, 0, 0, 0, time.UTC),
		constants.SongbirdID: time.Date(2100, time.January, 1, 0, 0, 0, 0, time.UTC),
		constants.FlareID:    time.Date(2100, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	ApricotPhase4DefaultTime = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)

	ApricotPhase5Times = map[uint32]time.Time{
		constants.CostonID:   time.Date(2100, time.January, 1, 0, 0, 0, 0, time.UTC),
		constants.SongbirdID: time.Date(2100, time.January, 1, 0, 0, 0, 0, time.UTC),
		constants.FlareID:    time.Date(2100, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	ApricotPhase5DefaultTime = time.Date(2000, time.December, 5, 5, 0, 0, 0, time.UTC)

	ApricotPhase4MinPChainHeight        = map[uint32]uint64{}
	ApricotPhase4DefaultMinPChainHeight = uint64(0)
)

func GetApricotPhase0Time(networkID uint32) time.Time {
	if upgradeTime, exists := ApricotPhase0Times[networkID]; exists {
		return upgradeTime
	}
	return ApricotPhase0DefaultTime
}

func GetApricotPhase1Time(networkID uint32) time.Time {
	if upgradeTime, exists := ApricotPhase1Times[networkID]; exists {
		return upgradeTime
	}
	return ApricotPhase1DefaultTime
}

func GetApricotPhase2Time(networkID uint32) time.Time {
	if upgradeTime, exists := ApricotPhase2Times[networkID]; exists {
		return upgradeTime
	}
	return ApricotPhase2DefaultTime
}

func GetApricotPhase3Time(networkID uint32) time.Time {
	if upgradeTime, exists := ApricotPhase3Times[networkID]; exists {
		return upgradeTime
	}
	return ApricotPhase3DefaultTime
}

func GetApricotPhase4Time(networkID uint32) time.Time {
	if upgradeTime, exists := ApricotPhase4Times[networkID]; exists {
		return upgradeTime
	}
	return ApricotPhase4DefaultTime
}

func GetApricotPhase4MinPChainHeight(networkID uint32) uint64 {
	if minHeight, exists := ApricotPhase4MinPChainHeight[networkID]; exists {
		return minHeight
	}
	return ApricotPhase4DefaultMinPChainHeight
}

func GetApricotPhase5Time(networkID uint32) time.Time {
	if upgradeTime, exists := ApricotPhase5Times[networkID]; exists {
		return upgradeTime
	}
	return ApricotPhase5DefaultTime
}

func GetCompatibility(networkID uint32) Compatibility {
	return NewCompatibility(
		CurrentApp,
		MinimumCompatibleVersion,
		GetApricotPhase5Time(networkID),
		PrevMinimumCompatibleVersion,
		MinimumUnmaskedVersion,
		GetApricotPhase0Time(networkID),
		PrevMinimumUnmaskedVersion,
	)
}
