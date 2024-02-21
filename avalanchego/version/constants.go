// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package version

import (
	"time"

	"github.com/flare-foundation/flare/utils/constants"
)

// These are globals that describe network upgrades and node versions
var (
	// Flare versioning constants.
	Current                      = NewDefaultVersion(0, 6, 3)
	CurrentApp                   = NewDefaultApplication(constants.PlatformName, Current.Major(), Current.Minor(), Current.Patch())
	MinimumCompatibleVersion     = NewDefaultApplication(constants.PlatformName, 0, 6, 0)
	PrevMinimumCompatibleVersion = NewDefaultApplication(constants.PlatformName, 0, 5, 1)
	MinimumUnmaskedVersion       = NewDefaultApplication(constants.PlatformName, 0, 5, 1)
	PrevMinimumUnmaskedVersion   = NewDefaultApplication(constants.PlatformName, 0, 5, 1)

	VersionParser = NewDefaultApplicationParser()

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
		constants.CostonID:   time.Date(2022, time.February, 25, 14, 0, 0, 0, time.UTC),
		constants.SongbirdID: time.Date(2022, time.March, 7, 14, 0, 0, 0, time.UTC),
	}
	ApricotPhase3DefaultTime = time.Date(2022, time.February, 9, 15, 0, 0, 0, time.UTC)

	ApricotPhase4Times = map[uint32]time.Time{
		constants.CostonID:   time.Date(2022, time.February, 25, 15, 0, 0, 0, time.UTC),
		constants.SongbirdID: time.Date(2022, time.March, 7, 15, 0, 0, 0, time.UTC),
	}
	ApricotPhase4DefaultTime = time.Date(2022, time.February, 10, 15, 0, 0, 0, time.UTC)

	ApricotPhase5Times = map[uint32]time.Time{
		constants.CostonID:   time.Date(2022, time.February, 25, 16, 0, 0, 0, time.UTC),
		constants.SongbirdID: time.Date(2022, time.March, 7, 16, 0, 0, 0, time.UTC),
	}
	ApricotPhase5DefaultTime = time.Date(2022, time.February, 11, 15, 0, 0, 0, time.UTC)

	StateConnectorTimes = map[uint32]time.Time{
		constants.CostonID:   time.Date(2022, time.March, 28, 14, 0, 0, 0, time.UTC),
		constants.SongbirdID: time.Date(2022, time.March, 28, 14, 0, 0, 0, time.UTC),
	}
	StateConnectorDefaultTime = time.Date(2022, time.February, 12, 15, 0, 0, 0, time.UTC)

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
func GetStateConnectorTime(networkID uint32) time.Time {
	if upgradeTime, exists := StateConnectorTimes[networkID]; exists {
		return upgradeTime
	}
	return StateConnectorDefaultTime
}

func GetCompatibility(networkID uint32) Compatibility {
	return NewCompatibility(
		CurrentApp,
		MinimumCompatibleVersion,
		GetStateConnectorTime(networkID),
		PrevMinimumCompatibleVersion,
		MinimumUnmaskedVersion,
		GetApricotPhase3Time(networkID),
		PrevMinimumUnmaskedVersion,
	)
}
