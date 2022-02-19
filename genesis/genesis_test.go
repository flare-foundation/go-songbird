// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
	"github.com/flare-foundation/flare/utils/hashing"
	"github.com/flare-foundation/flare/utils/perms"
	"github.com/flare-foundation/flare/vms/platformvm"
)

func TestValidateConfig(t *testing.T) {
	tests := map[string]struct {
		networkID uint32
		config    *Config
		err       string
	}{
		"flare": {
			networkID: 1,
			config:    &FlareConfig,
		},
		"songbird": {
			networkID: 5,
			config:    &SongbirdConfig,
		},
		"coston": {
			networkID: 7,
			config:    &CostonConfig,
		},
		"local": {
			networkID: 12345,
			config:    &LocalConfig,
		},
		"flare (networkID mismatch)": {
			networkID: 9,
			config:    &FlareConfig,
			err:       "networkID 9 specified but genesis config contains networkID 1",
		},
		"invalid start time": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.StartTime = 999999999999999
				return &thisConfig
			}(),
			err: "start time cannot be in the future",
		},
		"empty C-Chain genesis": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.CChainGenesis = ""
				return &thisConfig
			}(),
			err: "C-Chain genesis cannot be empty",
		},
		"empty message": {
			networkID: 12345,
			config: func() *Config {
				thisConfig := LocalConfig
				thisConfig.Message = ""
				return &thisConfig
			}(),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			err := validateConfig(test.networkID, test.config)
			if len(test.err) > 0 {
				assert.Error(err)
				assert.Contains(err.Error(), test.err)
				return
			}
			assert.NoError(err)
		})
	}
}

var (
	customGenesisConfigJSON = `{
		"networkID": 9999,
		"allocations": [
			{
				"ethAddr": "0xb3d82b1367d362de99ab59a658165aff520cbd4d",
				"avaxAddr": "X-local1g65uqn6t77p656w64023nh8nd9updzmxyymev2",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			},
			{
				"ethAddr": "0xb3d82b1367d362de99ab59a658165aff520cbd4d",
				"avaxAddr": "X-local18jma8ppw3nhx5r4ap8clazz0dps7rv5u00z96u",
				"initialAmount": 300000000000000000,
				"unlockSchedule": [
					{
						"amount": 20000000000000000
					},
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			},
			{
				"ethAddr": "0xb3d82b1367d362de99ab59a658165aff520cbd4d",
				"avaxAddr": "X-local1ur873jhz9qnaqv5qthk5sn3e8nj3e0kmggalnu",
				"initialAmount": 10000000000000000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000,
						"locktime": 1633824000
					}
				]
			}
		],
		"startTime": 1599696000,
		"initialStakeDuration": 31536000,
		"initialStakeDurationOffset": 5400,
		"initialStakedFunds": [
			"X-local1g65uqn6t77p656w64023nh8nd9updzmxyymev2"
		],
		"initialStakers": [
			{
				"nodeID": "NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg",
				"rewardAddress": "X-local18jma8ppw3nhx5r4ap8clazz0dps7rv5u00z96u",
				"delegationFee": 1000000
			},
			{
				"nodeID": "NodeID-MFrZFVCXPv5iCn6M9K6XduxGTYp891xXZ",
				"rewardAddress": "X-local18jma8ppw3nhx5r4ap8clazz0dps7rv5u00z96u",
				"delegationFee": 500000
			},
			{
				"nodeID": "NodeID-NFBbbJ4qCmNaCzeW7sxErhvWqvEQMnYcN",
				"rewardAddress": "X-local18jma8ppw3nhx5r4ap8clazz0dps7rv5u00z96u",
				"delegationFee": 250000
			},
			{
				"nodeID": "NodeID-GWPcbFJZFfZreETSoWjPimr846mXEKCtu",
				"rewardAddress": "X-local18jma8ppw3nhx5r4ap8clazz0dps7rv5u00z96u",
				"delegationFee": 125000
			},
			{
				"nodeID": "NodeID-P7oB2McjBGgW2NXXWVYjV8JEDFoW9xDE5",
				"rewardAddress": "X-local18jma8ppw3nhx5r4ap8clazz0dps7rv5u00z96u",
				"delegationFee": 62500
			}
		],
		"cChainGenesis": "{\"config\":{\"chainId\":1337,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x5f5e100\",\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"0100000000000000000000000000000000000000\":{\"code\":\"0x7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c80631e010439146042578063b6510bb314606e575b600080fd5b605c60048036036020811015605657600080fd5b503560b1565b60408051918252519081900360200190f35b818015607957600080fd5b5060af60048036036080811015608e57600080fd5b506001600160a01b03813516906020810135906040810135906060013560b6565b005b30cd90565b836001600160a01b031681836108fc8690811502906040516000604051808303818888878c8acf9550505050505015801560f4573d6000803e3d6000fd5b505050505056fea26469706673582212201eebce970fe3f5cb96bf8ac6ba5f5c133fc2908ae3dcd51082cfee8f583429d064736f6c634300060a0033\",\"balance\":\"0x0\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
		"message": "{{ fun_quote }}"
	}`
	invalidGenesisConfigJSON = `{
		"networkID": 9999}}}}
	}`
)

func TestGenesisFromFile(t *testing.T) {
	tests := map[string]struct {
		networkID       uint32
		customConfig    string
		missingFilepath string
		err             string
		expected        string
	}{
		"flare": {
			networkID:    constants.FlareID,
			customConfig: customGenesisConfigJSON,
			err:          "cannot override genesis config for standard network flare (1)",
		},
		"songbird": {
			networkID:    constants.SongbirdID,
			customConfig: customGenesisConfigJSON,
			err:          "cannot override genesis config for standard network songbird (5)",
		},
		"coston": {
			networkID:    constants.CostonID,
			customConfig: customGenesisConfigJSON,
			err:          "cannot override genesis config for standard network coston (7)",
		},
		"songbird (with custom specified)": {
			networkID:    constants.SongbirdID,
			customConfig: localGenesisConfigJSON, // won't load
			err:          "cannot override genesis config for standard network songbird (5)",
		},
		"local": {
			networkID:    constants.LocalID,
			customConfig: customGenesisConfigJSON,
			err:          "cannot override genesis config for standard network local (12345)",
		},
		"local (with custom specified)": {
			networkID:    constants.LocalID,
			customConfig: customGenesisConfigJSON,
			err:          "cannot override genesis config for standard network local (12345)",
		},
		"custom": {
			networkID:    9999,
			customConfig: customGenesisConfigJSON,
			expected:     "11a5e4f9616887080e10f7d0ad0638b717e00ff7f390f29d3e72c0da995d05bd",
		},
		"custom (networkID mismatch)": {
			networkID:    9999,
			customConfig: localGenesisConfigJSON,
			err:          "networkID 9999 specified but genesis config contains networkID 12345",
		},
		"custom (invalid format)": {
			networkID:    9999,
			customConfig: invalidGenesisConfigJSON,
			err:          "unable to load provided genesis config",
		},
		"custom (missing filepath)": {
			networkID:       9999,
			missingFilepath: "missing.json",
			err:             "unable to load provided genesis config",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// test loading of genesis from file

			assert := assert.New(t)
			var customFile string
			if len(test.customConfig) > 0 {
				customFile = filepath.Join(t.TempDir(), "config.json")
				assert.NoError(perms.WriteFile(customFile, []byte(test.customConfig), perms.ReadWrite))
			}

			if len(test.missingFilepath) > 0 {
				customFile = test.missingFilepath
			}

			genesisBytes, _, err := FromFile(test.networkID, customFile)
			if len(test.err) > 0 {
				assert.Error(err)
				assert.Contains(err.Error(), test.err)
				return
			}
			assert.NoError(err)

			genesisHash := fmt.Sprintf("%x", hashing.ComputeHash256(genesisBytes))
			assert.Equal(test.expected, genesisHash, "genesis hash mismatch")

			genesis := platformvm.Genesis{}
			_, err = platformvm.GenesisCodec.Unmarshal(genesisBytes, &genesis)
			assert.NoError(err)
		})
	}
}

func TestGenesisFromFlag(t *testing.T) {
	tests := map[string]struct {
		networkID    uint32
		customConfig string
		err          string
		expected     string
	}{
		"flare": {
			networkID: constants.FlareID,
			err:       "cannot override genesis config for standard network flare (1)",
		},
		"songbird": {
			networkID: constants.SongbirdID,
			err:       "cannot override genesis config for standard network songbird (5)",
		},
		"coston": {
			networkID: constants.CostonID,
			err:       "cannot override genesis config for standard network coston (7)",
		},
		"local": {
			networkID: constants.LocalID,
			err:       "cannot override genesis config for standard network local (12345)",
		},
		"local (with custom specified)": {
			networkID:    constants.LocalID,
			customConfig: customGenesisConfigJSON,
			err:          "cannot override genesis config for standard network local (12345)",
		},
		"custom": {
			networkID:    9999,
			customConfig: customGenesisConfigJSON,
			expected:     "11a5e4f9616887080e10f7d0ad0638b717e00ff7f390f29d3e72c0da995d05bd",
		},
		"custom (networkID mismatch)": {
			networkID:    9999,
			customConfig: localGenesisConfigJSON,
			err:          "networkID 9999 specified but genesis config contains networkID 12345",
		},
		"custom (invalid format)": {
			networkID:    9999,
			customConfig: invalidGenesisConfigJSON,
			err:          "unable to load genesis content from flag",
		},
		"custom (missing content)": {
			networkID: 9999,
			err:       "unable to load genesis content from flag",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// test loading of genesis content from flag/env-var

			assert := assert.New(t)
			var genBytes []byte
			if len(test.customConfig) == 0 {
				// try loading a default config
				var err error
				switch test.networkID {
				case constants.FlareID:
					genBytes, err = json.Marshal(&FlareConfig)
					assert.NoError(err)
				case constants.SongbirdID:
					genBytes, err = json.Marshal(&SongbirdConfig)
					assert.NoError(err)
				case constants.CostonID:
					genBytes, err = json.Marshal(&CostonConfig)
					assert.NoError(err)
				case constants.LocalID:
					genBytes, err = json.Marshal(&LocalConfig)
					assert.NoError(err)
				default:
					genBytes = make([]byte, 0)
				}
			} else {
				genBytes = []byte(test.customConfig)
			}
			content := base64.StdEncoding.EncodeToString(genBytes)

			genesisBytes, _, err := FromFlag(test.networkID, content)
			if len(test.err) > 0 {
				assert.Error(err)
				assert.Contains(err.Error(), test.err)
				return
			}
			assert.NoError(err)

			genesisHash := fmt.Sprintf("%x", hashing.ComputeHash256(genesisBytes))
			assert.Equal(test.expected, genesisHash, "genesis hash mismatch")

			genesis := platformvm.Genesis{}
			_, err = platformvm.GenesisCodec.Unmarshal(genesisBytes, &genesis)
			assert.NoError(err)
		})
	}
}

func TestVMGenesis(t *testing.T) {
	type vmTest struct {
		vmID       ids.ID
		expectedID string
	}
	tests := []struct {
		networkID uint32
		vmTest    []vmTest
	}{
		{
			networkID: constants.FlareID,
			vmTest: []vmTest{
				{
					vmID:       constants.AVMID,
					expectedID: "kuZe3hRrShPqeGbHag3ffGVNUTeca2TgFmbDhKgB9gPyCuvaq",
				},
				{
					vmID:       constants.EVMID,
					expectedID: "2K181beJgGqi3z1N8gQC6Tq1ffejTiH3HZRGRYukKjyFDXspwp",
				},
			},
		},
		{
			networkID: constants.SongbirdID,
			vmTest: []vmTest{
				{
					vmID:       constants.AVMID,
					expectedID: "7xKYhEvYuUekwDxozgEiMPufzJ3jJPypKbGE8ny6KL84z4RKB",
				},
				{
					vmID:       constants.EVMID,
					expectedID: "erCt5pSo5d4bM8fMrsB2dRM54PGssDAVqRg1jHedQzr6ayLiq",
				},
			},
		},
		{
			networkID: constants.CostonID,
			vmTest: []vmTest{
				{
					vmID:       constants.AVMID,
					expectedID: "8S5vg8ycMqULMMq2CzrRh3qqbFWjs6a35H8UwNAaS83v9Ynxa",
				},
				{
					vmID:       constants.EVMID,
					expectedID: "2wcJd3HZL4EAoQDfGhyeHGJU55wizganSzD9EvdDLJX3B5u5Be",
				},
			},
		},
		{
			networkID: constants.LocalID,
			vmTest: []vmTest{
				{
					vmID:       constants.AVMID,
					expectedID: "ALRkp1tuy7ErVkWuEWFLVd657JAULWDDyQkQBkLKVE94jCaNu",
				},
				{
					vmID:       constants.EVMID,
					expectedID: "qNigBiF2gjcLWf8Zguzi1ztr2nJ87Z7bWwbDzWBJvuwuhvBkX",
				},
			},
		},
	}

	for _, test := range tests {
		for _, vmTest := range test.vmTest {
			name := fmt.Sprintf("%s-%s",
				constants.NetworkIDToNetworkName[test.networkID],
				vmTest.vmID,
			)
			t.Run(name, func(t *testing.T) {
				assert := assert.New(t)

				config := GetConfig(test.networkID)
				genesisBytes, _, err := FromConfig(config)
				assert.NoError(err)

				genesisTx, err := VMGenesis(genesisBytes, vmTest.vmID)
				assert.NoError(err)

				assert.Equal(
					vmTest.expectedID,
					genesisTx.ID().String(),
					"%s genesisID with networkID %d mismatch",
					vmTest.vmID,
					test.networkID,
				)
			})
		}
	}
}

func TestAVAXAssetID(t *testing.T) {
	tests := []struct {
		networkID  uint32
		expectedID string
	}{
		{
			networkID:  constants.FlareID,
			expectedID: "foMCFvzKECiGVJmmkAEHm9Vt43hYjuxreiNX5PfqfecaVsZBT",
		},
		{
			networkID:  constants.SongbirdID,
			expectedID: "1S3PSi4VsVpD8iK2vdykuajxVeuCV2xhjPSkQ4K88mqWGozMP",
		},
		{
			networkID:  constants.CostonID,
			expectedID: "sMWvCcweHFoG4SE1SSVtiwosqezUW8tAA77t7BhAS7RUh6ms4",
		},
		{
			networkID:  constants.LocalID,
			expectedID: "2RULRJVXVpQNAsV3sBpy4G8LWH1LN3z5Adokv5bVtnZmsBQDCX",
		},
	}

	for _, test := range tests {
		t.Run(constants.NetworkIDToNetworkName[test.networkID], func(t *testing.T) {
			assert := assert.New(t)

			config := GetConfig(test.networkID)
			_, avaxAssetID, err := FromConfig(config)
			assert.NoError(err)

			assert.Equal(
				test.expectedID,
				avaxAssetID.String(),
				"AVAX assetID with networkID %d mismatch",
				test.networkID,
			)
		})
	}
}
