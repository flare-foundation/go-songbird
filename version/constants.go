// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package version

import (
	"fmt"
	"time"

	"github.com/flare-foundation/flare/utils/constants"
)

var (
	String                       string // Printed when CLI arg --version is used
	GitCommit                    string // Set in the build script (i.e. at compile time)
	Current                      = NewDefaultVersion(0, 1, 2)
	CurrentApp                   = NewDefaultApplication(constants.PlatformName, Current.Major(), Current.Minor(), Current.Patch())
	MinimumCompatibleVersion     = NewDefaultApplication(constants.PlatformName, 0, 1, 0)
	PrevMinimumCompatibleVersion = NewDefaultApplication(constants.PlatformName, 0, 0, 1)
	MinimumUnmaskedVersion       = NewDefaultApplication(constants.PlatformName, 0, 1, 0)
	PrevMinimumUnmaskedVersion   = NewDefaultApplication(constants.PlatformName, 0, 0, 1)
	VersionParser                = NewDefaultApplicationParser()

	CurrentDatabase = DatabaseVersion1_4_5
	PrevDatabase    = DatabaseVersion1_0_0

	DatabaseVersion1_4_5 = NewDefaultVersion(1, 4, 5)
	DatabaseVersion1_0_0 = NewDefaultVersion(1, 0, 0)

	ApricotPhase0Times       = map[uint32]time.Time{}
	ApricotPhase0DefaultTime = time.Date(2021, time.January, 0, 0, 0, 0, 0, time.UTC)

	ApricotPhase1Times       = map[uint32]time.Time{}
	ApricotPhase1DefaultTime = time.Date(2021, time.January, 0, 0, 0, 0, 0, time.UTC)

	ApricotPhase2Times       = map[uint32]time.Time{}
	ApricotPhase2DefaultTime = time.Date(2021, time.January, 0, 0, 0, 0, 0, time.UTC)

	ApricotPhase3Times       = map[uint32]time.Time{}
	ApricotPhase3DefaultTime = time.Date(2021, time.January, 0, 0, 0, 0, 0, time.UTC)
)

func init() {
	format := "%s [database=%s"
	args := []interface{}{
		CurrentApp,
		CurrentDatabase,
	}
	if GitCommit != "" {
		format += ", commit=%s"
		args = append(args, GitCommit)
	}
	format += "]\n"
	String = fmt.Sprintf(format, args...)
}

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

func GetCompatibility(networkID uint32) Compatibility {
	return NewCompatibility(
		CurrentApp,
		MinimumCompatibleVersion,
		GetApricotPhase3Time(networkID),
		PrevMinimumCompatibleVersion,
		MinimumUnmaskedVersion,
		GetApricotPhase0Time(networkID),
		PrevMinimumUnmaskedVersion,
	)
}
