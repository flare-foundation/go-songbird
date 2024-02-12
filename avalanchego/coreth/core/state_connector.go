// (c) 2021, Flare Networks Limited. All rights reserved.
// Please see the file LICENSE for licensing terms.

package core

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/coreth/core/vm"
	"github.com/flare-foundation/flare/coreth/params"
)

const (
	localAttestorEnv  = "SC_LOCAL_ATTESTATORS_SONGBIRD"
	forkingEnabledEnv = "SC_FORKING_ENABLED_SONGBIRD"
)

var (
	costonChainID   = new(big.Int).SetUint64(16) // https://github.com/ethereum-lists/chains/blob/master/_data/chains/eip155-16.json
	songbirdChainID = new(big.Int).SetUint64(19) // https://github.com/ethereum-lists/chains/blob/master/_data/chains/eip155-19.json

	costonActivationTime   = big.NewInt(time.Date(2022, time.February, 25, 17, 0, 0, 0, time.UTC).Unix())
	songbirdActivationTime = big.NewInt(time.Date(2022, time.March, 28, 14, 0, 0, 0, time.UTC).Unix())

	costonOct22ForkTime   = big.NewInt(time.Date(2022, time.October, 6, 15, 0, 0, 0, time.UTC).Unix())
	songbirdOct22ForkTime = big.NewInt(time.Date(2022, time.October, 19, 15, 0, 0, 0, time.UTC).Unix())
)

type AttestationVotes struct {
	reachedMajority    bool
	majorityDecision   string
	majorityAttestors  []common.Address
	divergentAttestors []common.Address
	abstainedAttestors []common.Address
}

func GetStateConnectorActivated(chainID *big.Int, blockTime *big.Int) bool {
	switch {
	case chainID.Cmp(songbirdChainID) == 0:
		return blockTime.Cmp(songbirdActivationTime) >= 0
	case chainID.Cmp(costonChainID) == 0:
		return blockTime.Cmp(costonActivationTime) >= 0
	default:
		return true
	}
}

func GetStateConnectorContract(chainID *big.Int, blockTime *big.Int) common.Address {
	switch {
	case chainID.Cmp(songbirdChainID) == 0:
		switch {
		case blockTime.Cmp(songbirdOct22ForkTime) > 0:
			return common.HexToAddress("0x0c13aDA1C7143Cf0a0795FFaB93eEBb6FAD6e4e3")
		default:
			return common.HexToAddress("0x3A1b3220527aBA427d1e13e4b4c48c31460B4d91")
		}
	case chainID.Cmp(costonChainID) == 0:
		switch {
		case blockTime.Cmp(costonOct22ForkTime) > 0:
			return common.HexToAddress("0x0c13aDA1C7143Cf0a0795FFaB93eEBb6FAD6e4e3")
		default:
			return common.HexToAddress("0x947c76694491d3fD67a73688003c4d36C8780A97")
		}
	default:
		return common.HexToAddress("0x1000000000000000000000000000000000000001")
	}
}

func GetStateConnectorCoinbaseSignalAddr(chainID *big.Int, blockTime *big.Int) common.Address {
	switch {
	case chainID.Cmp(songbirdChainID) == 0:
		switch {
		case blockTime.Cmp(songbirdOct22ForkTime) > 0:
			return common.HexToAddress("0x00000000000000000000000000000000000DEaD1")
		default:
			return common.HexToAddress("0x000000000000000000000000000000000000dEaD")
		}
	case chainID.Cmp(costonChainID) == 0:
		switch {
		case blockTime.Cmp(costonOct22ForkTime) > 0:
			return common.HexToAddress("0x00000000000000000000000000000000000DEaD1")
		default:
			return common.HexToAddress("0x000000000000000000000000000000000000dEaD")
		}
	default:
		return common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	}
}

func SubmitAttestationSelector(chainID *big.Int, blockTime *big.Int) []byte {
	switch {
	default:
		return []byte{0xcf, 0xd1, 0xfd, 0xad}
	}
}

func GetAttestationSelector(chainID *big.Int, blockTime *big.Int) []byte {
	switch {
	default:
		return []byte{0x29, 0xbe, 0x4d, 0xb2}
	}
}

func FinaliseRoundSelector(chainID *big.Int, blockTime *big.Int) []byte {
	switch {
	default:
		return []byte{0xea, 0xeb, 0xf6, 0xd3}
	}
}

// The default attestation providers for the state connector will be drawn from the top weighted/performing FTSOs.
func GetDefaultAttestors(chainID *big.Int, blockTime *big.Int) []common.Address {
	switch {
	case chainID.Cmp(songbirdChainID) == 0:
		switch {
		case blockTime.Cmp(submitterContractActivationTimeSongbird) > 0:
			return []common.Address{
				common.HexToAddress("0x2D3e7e4b19bDc920fd9C57BD3072A31F5a59FeC8"),
				common.HexToAddress("0x442DD539Fe78D43A1a9358FF3460CfE63e2bC9CC"),
				common.HexToAddress("0x49893c5Dfc035F4eE4E46faC014f6D4bC80F7f92"),
				common.HexToAddress("0x5D2f75392DdDa69a2818021dd6a64937904c8352"),
				common.HexToAddress("0x6455dC38fdF739b6fE021b30C7D9672C1c6DEb5c"),
				common.HexToAddress("0x808441Ec3Fa1721330226E69527Bc160D8d9386a"),
				common.HexToAddress("0x823B0f5c7758E9d3bE55bA1EA840E29ccd5D5CcB"),
				common.HexToAddress("0x85016969b9eBDB8977975a4743c9FCEeabCEAf8A"),
				common.HexToAddress("0x8A3D627D86A81F5D21683F4963565C63DB5e1309"),
			}
		case blockTime.Cmp(songbirdOct22ForkTime) > 0:
			return []common.Address{
				common.HexToAddress("0x2D3e7e4b19bDc920fd9C57BD3072A31F5a59FeC8"),
				common.HexToAddress("0x442DD539Fe78D43A1a9358FF3460CfE63e2bC9CC"),
				common.HexToAddress("0x49893c5Dfc035F4eE4E46faC014f6D4bC80F7f92"),
				common.HexToAddress("0x5D2f75392DdDa69a2818021dd6a64937904c8352"),
				common.HexToAddress("0x6455dC38fdF739b6fE021b30C7D9672C1c6DEb5c"),
				common.HexToAddress("0x808441Ec3Fa1721330226E69527Bc160D8d9386a"),
				common.HexToAddress("0x823B0f5c7758E9d3bE55bA1EA840E29ccd5D5CcB"),
				common.HexToAddress("0x85016969b9eBDB8977975a4743c9FCEeabCEAf8A"),
				common.HexToAddress("0x8A3D627D86A81F5D21683F4963565C63DB5e1309"),
			}
		default:
			return []common.Address{
				common.HexToAddress("0x0c19f3B4927abFc596353B0f9Ddad5D817736F70"),
			}
		}
	case chainID.Cmp(costonChainID) == 0:
		switch {
		case blockTime.Cmp(costonOct22ForkTime) > 0:
			return []common.Address{
				common.HexToAddress("0x30e4b4542b4aAf615838B113f14c46dE1469212e"),
				common.HexToAddress("0x3519E14183252794aaA52aA824f34482ef44cE1d"),
				common.HexToAddress("0xb445857476181ec378Ec453ab3d122183CfC3b78"),
				common.HexToAddress("0x6D755cd7A61A9DCFc96FaE0f927C3a73bE986ce4"),
				common.HexToAddress("0xdC0fD24846303D58d2D66AA8820be2685735dBd2"),
				common.HexToAddress("0x3F52c41c0500a4f018A38c9f8273b254aD7e2FCc"),
				common.HexToAddress("0xdA6d6aA9F1f770c279c5DA0C71f4DC1142A70d5D"),
				common.HexToAddress("0x3d895D00d2802120D39d4D2554F7ef09d6845E99"),
				common.HexToAddress("0xc36141CFBe5Af6eB2F8b21550Ccd457DA7FaF3C6"),
			}
		default:
			return []common.Address{
				common.HexToAddress("0x3a6e101103ec3d9267d08f484a6b70e1440a8255"),
			}
		}
	default:
		return []common.Address{}
	}
}

func GetLocalAttestors() []common.Address {
	var localAttestors []common.Address
	localAttestorList := os.Getenv(localAttestorEnv)
	if localAttestorList != "" {
		localAttestorEntries := strings.Split(localAttestorList, ",")
		for _, localAttestorEntry := range localAttestorEntries {
			localAttestors = append(localAttestors, common.HexToAddress(localAttestorEntry))
		}
	}
	return localAttestors
}

func (st *StateTransition) GetAttestation(attestor common.Address, instructions []byte) (string, error) {
	merkleRootHash, _, err := st.evm.Call(vm.AccountRef(attestor), st.to(), instructions, params.TxGas, big.NewInt(0))
	return hex.EncodeToString(merkleRootHash), err
}

func (st *StateTransition) GetAttestations(attestors []common.Address, instructions []byte) (AttestationVotes, int, map[string][]common.Address) {
	var attestationVotes AttestationVotes
	hashFrequencies := make(map[string][]common.Address)
	for i, a := range attestors {
		h, err := st.GetAttestation(a, instructions)
		if err != nil {
			attestationVotes.abstainedAttestors = append(attestationVotes.abstainedAttestors, a)
		}
		hashFrequencies[h] = append(hashFrequencies[h], attestors[i])
	}
	return attestationVotes, len(attestors), hashFrequencies
}

func CountAttestations(attestationVotes AttestationVotes, numAttestors int, hashFrequencies map[string][]common.Address) AttestationVotes {
	// Find the plurality
	var pluralityNum int
	var pluralityKey string
	for key, val := range hashFrequencies {
		if len(val) > pluralityNum && len(key) > 0 {
			pluralityNum = len(val)
			pluralityKey = key
		}
	}
	if pluralityNum > numAttestors/2 {
		attestationVotes.reachedMajority = true
		attestationVotes.majorityDecision = pluralityKey
		attestationVotes.majorityAttestors = hashFrequencies[pluralityKey]
	}
	for key, val := range hashFrequencies {
		if key != pluralityKey {
			attestationVotes.divergentAttestors = append(attestationVotes.divergentAttestors, val...)
		}
	}
	return attestationVotes
}

func (st *StateTransition) FinalisePreviousRound(chainID *big.Int, timestamp *big.Int, currentRoundNumber []byte) error {
	getAttestationSelector := GetAttestationSelector(chainID, timestamp)
	instructions := append(getAttestationSelector[:], currentRoundNumber[:]...)
	defaultAttestors := GetDefaultAttestors(chainID, timestamp)
	defaultAttestationVotes := CountAttestations(st.GetAttestations(defaultAttestors, instructions))
	localAttestors := GetLocalAttestors()
	finalityReached := defaultAttestationVotes.reachedMajority
	if len(localAttestors) > 0 {
		localAttestationVotes := CountAttestations(st.GetAttestations(localAttestors, instructions))
		if finalityReached && defaultAttestationVotes.majorityDecision != localAttestationVotes.majorityDecision && os.Getenv(forkingEnabledEnv) == "1" {
			// Fork this node now from the default path
			return fmt.Errorf(
				"default state connector decision (%s) does not match this node's local state connector decision (%s), forking node",
				defaultAttestationVotes.majorityDecision,
				localAttestationVotes.majorityDecision,
			)
		}
	}
	if finalityReached {
		// Finalise defaultAttestationVotes.majorityDecision
		finaliseRoundSelector := FinaliseRoundSelector(chainID, timestamp)
		finalisedData := append(finaliseRoundSelector[:], currentRoundNumber[:]...)
		merkleRootHashBytes, err := hex.DecodeString(defaultAttestationVotes.majorityDecision)
		if err != nil {
			return err
		}
		finalisedData = append(finalisedData[:], merkleRootHashBytes[:]...)
		coinbaseSignal := GetStateConnectorCoinbaseSignalAddr(chainID, timestamp)
		originalCoinbase := st.evm.Context.Coinbase
		defer func() {
			st.evm.Context.Coinbase = originalCoinbase
		}()
		// Setting msg.sender = block.coinbase and block.coinbase = SIGNAL_COINBASE signals to the EVM to finalise this round
		// See: https://gitlab.com/flarenetwork/flare-smart-contracts/-/blob/57ac7259f1708832201b774fc3445e0fbfb94ef4/contracts/genesis/implementation/StateConnector.sol#L132
		st.evm.Context.Coinbase = coinbaseSignal
		// In order to break the State Connector's signalling mechanism, one would have to both:
		// 		1) Change the Flare validator code to enable them to control the block.coinbase variable. This is mitigated in state_transition.go
		//				by this check: burnAddress == common.HexToAddress("0x0100000000000000000000000000000000000000") on line 373, which occurs
		//				right before st.FinalisePreviousRound(chainID, timestamp, st.data[4:36]) is called.
		//		2) Know the private key to the address 0x00000000000000000000000000000000000DEaD1 in order to become msg.sender.
		_, _, err = st.evm.Call(vm.AccountRef(coinbaseSignal), st.to(), finalisedData, st.evm.Context.GasLimit, big.NewInt(0))
		if err != nil {
			return err
		}
	}
	return nil
}
