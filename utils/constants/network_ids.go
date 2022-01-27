// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package constants

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/flare-foundation/flare/ids"
)

// Const variables to be exported
const (
	FlareID    uint32 = 1
	SongbirdID uint32 = 5
	CostonID   uint32 = 7
	TestingID  uint32 = 10
	LocalID    uint32 = 12345

	FlareName    = "flare"
	SongbirdName = "songbird"
	CostonName   = "coston"
	TestingName  = "testing"
	LocalName    = "local"

	FlareHRP    = "flare"
	SongbirdHRP = "songbird"
	CostonHRP   = "coston"
	TestingHRP  = "testing"
	LocalHRP    = "local"
	FallbackHRP = "custom"
)

// Variables to be exported
var (
	PrimaryNetworkID = ids.Empty
	PlatformChainID  = ids.Empty

	NetworkIDToNetworkName = map[uint32]string{
		FlareID:    FlareName,
		SongbirdID: SongbirdName,
		CostonID:   CostonName,
		TestingID:  TestingName,
		LocalID:    LocalName,
	}
	NetworkNameToNetworkID = map[string]uint32{
		FlareName:    FlareID,
		SongbirdName: SongbirdID,
		CostonName:   CostonID,
		TestingName:  TestingID,
		LocalName:    LocalID,
	}

	NetworkIDToHRP = map[uint32]string{
		FlareID:    FlareHRP,
		SongbirdID: SongbirdHRP,
		CostonID:   CostonHRP,
		TestingID:  TestingHRP,
		LocalID:    LocalHRP,
	}
	NetworkHRPToNetworkID = map[string]uint32{
		FlareHRP:    FlareID,
		SongbirdHRP: SongbirdID,
		CostonHRP:   CostonID,
		TestingHRP:  TestingID,
		LocalHRP:    LocalID,
	}

	ValidNetworkPrefix = "network-"
)

// GetHRP returns the Human-Readable-Part of bech32 addresses for a networkID
func GetHRP(networkID uint32) string {
	if hrp, ok := NetworkIDToHRP[networkID]; ok {
		return hrp
	}
	return FallbackHRP
}

// NetworkName returns a human readable name for the network with
// ID [networkID]
func NetworkName(networkID uint32) string {
	if name, exists := NetworkIDToNetworkName[networkID]; exists {
		return name
	}
	return fmt.Sprintf("network-%d", networkID)
}

// NetworkID returns the ID of the network with name [networkName]
func NetworkID(networkName string) (uint32, error) {
	networkName = strings.ToLower(networkName)
	if id, exists := NetworkNameToNetworkID[networkName]; exists {
		return id, nil
	}

	idStr := networkName
	if strings.HasPrefix(networkName, ValidNetworkPrefix) {
		idStr = networkName[len(ValidNetworkPrefix):]
	}
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %q as a network name", networkName)
	}
	return uint32(id), nil
}
