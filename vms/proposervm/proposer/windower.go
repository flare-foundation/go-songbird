// (c) 2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposer

import (
	"fmt"
	"github.com/flare-foundation/flare/snow/engine/snowman/block"
	"github.com/flare-foundation/flare/utils/subprocess"
	"github.com/flare-foundation/flare/vms/rpcchainvm"
	"github.com/flare-foundation/flare/vms/validatorvm"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"io/ioutil"
	"log"
	"sort"
	"time"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/validators"
	"github.com/flare-foundation/flare/utils/math"
	"github.com/flare-foundation/flare/utils/sampler"
	"github.com/flare-foundation/flare/utils/wrappers"
)

// Proposer list constants
const (
	MaxWindows     = 6
	WindowDuration = 5 * time.Second
	MaxDelay       = MaxWindows * WindowDuration
)

var _ Windower = &windower{}

type Windower interface {
	Delay(
		chainHeight,
		pChainHeight uint64,
		validatorID ids.ShortID,
		hash ids.ID,
	) (time.Duration, error)
	GetValidators(hash []byte) (map[string]float64, error)
}

// windower interfaces with P-Chain and it is responsible for calculating the
// delay for the block submission window of a given validator
type windower struct {
	state        validators.State
	subnetID     ids.ID
	chainSource  uint64
	sampler      sampler.WeightedWithoutReplacement
	vmValidator  *validatorvm.ValidatorVM
	valClient    *rpcchainvm.ValidatorsClient
	valInterface *block.ValidatorVMInterface
}

func New(state validators.State, subnetID, chainID ids.ID, vmValidator *validatorvm.ValidatorVM) Windower {
	fmt.Println("windower new called......")
	w := wrappers.Packer{Bytes: chainID[:]}
	fmt.Println(rpcchainvm.PluginMap) // todo how about create an exportable client as well? along with PluginMap
	//valClient := rpcchainvm.PluginMap["validators"].(*rpcchainvm.PluginValidator).ValVM.(block.ValidatorVMInterface).(*rpcchainvm.ValidatorsClient)
	//valClient := dispense(rpcchainvm.PluginMap)
	return &windower{
		state:       state,
		subnetID:    subnetID,
		chainSource: w.UnpackLong(),
		sampler:     sampler.NewDeterministicWeightedWithoutReplacement(),
		vmValidator: vmValidator,
		valClient:   rpcchainvm.GlobalValidatorClient,
	}
}

func dispense(pluginmap map[string]plugin.Plugin) *rpcchainvm.ValidatorsClient {
	// Ignore warning from launching an executable with a variable command
	// because the command is a controlled and required input
	PluginMap := pluginmap
	config := &plugin.ClientConfig{
		HandshakeConfig: rpcchainvm.Handshake,
		Plugins:         PluginMap,
		Cmd:             subprocess.New("/Users/default/go/src/github.com/flare/build/plugins/evm"), //f.Path
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC,
			plugin.ProtocolGRPC,
		},
		// We kill this client by calling kill() when the chain running this VM
		// shuts down. However, there are some cases where the VM's Shutdown
		// method is not called. Namely, if:
		// 1) The node shuts down after the client is created but before the
		//    chain is registered with the message router.
		// 2) The chain doesn't handle a shutdown message before the node times
		//    out on the chain's shutdown and dies, leaving the shutdown message
		//    unhandled.
		// We set managed to true so that we can call plugin.CleanupClients on
		// node shutdown to ensure every plugin subprocess is killed.
		Managed: true,
	}

	log.SetOutput(ioutil.Discard)
	config.Stderr = ioutil.Discard
	config.Logger = hclog.New(&hclog.LoggerOptions{
		Output: ioutil.Discard,
	})

	client := plugin.NewClient(config)

	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("rpcClient nil with error: ", err.Error())
		client.Kill()
	}
	fmt.Println("rpclient: ", rpcClient)
	raw1, err := rpcClient.Dispense("validators") //validators
	if err != nil {
		client.Kill()
	}
	valVM, ok := raw1.(*rpcchainvm.ValidatorsClient)
	if !ok {
		fmt.Println("cant convert")
		client.Kill()
	}
	fmt.Println("valVM: ", valVM)
	valVM.SetProcess(client)
	//idString := "11111111111111111111111111111111LpoYY"
	//
	//id, err := ids.ToID([]byte(idString))
	//if err != nil {
	//	fmt.Println("error in converting to id", err.Error())
	//}
	//testRet, err := valVM.GetValidators(id)
	//if err != nil {
	//	fmt.Println("return from GetValidators test: ", testRet)
	//	fmt.Println(err.Error())
	//}
	return valVM
}

func (w *windower) Delay(chainHeight, pChainHeight uint64, validatorID ids.ShortID, hash ids.ID) (time.Duration, error) {
	//todo take hash of parent block , remove pChainHeight
	//todo make separate proto file!
	// todo blockID
	//bID, _ := w.vmValidator.VM.State.GetLastAccepted() //todo do we really need to use the VM here or can we do without it so that we can get rid of the circular dependency?
	//(b.ParentID().([]byte))
	fmt.Println("Inside windower Delay()")
	fmt.Println(w.valClient)
	//validatorsMapNew, err := w.vmValidator.GetValidators(hash)
	validatorsMapNew, err := w.valClient.GetValidators(hash)
	fmt.Println("validatorsMapNew: ", validatorsMapNew)
	if validatorID == ids.ShortEmpty {
		return MaxDelay, nil
	}

	// get the validator set by the p-chain height
	validatorsMap, err := w.state.GetValidatorSet(pChainHeight, w.subnetID)
	//todo use the newly made Flare api call instead of the above function. (todo Do it just like the Postman call?)
	//todo use the chainHeight and pass it to create header in the flare call.
	if err != nil {
		return 0, err
	}

	// convert the map of validators to a slice
	validators := make(validatorsSlice, 0, len(validatorsMap))
	weight := uint64(0)
	for k, v := range validatorsMap {
		validators = append(validators, validatorData{
			id:     k,
			weight: v,
		})
		newWeight, err := math.Add64(weight, v)
		if err != nil {
			return 0, err
		}
		weight = newWeight
	}
	// New validators from coreth
	validators = nil
	for id, u := range validatorsMapNew {
		//longID, err := ids.ToID(id)
		//shortID, err := ids.ToShortID(id)
		if err != nil {
			continue
		}
		validators = append(validators, validatorData{
			id:        id, //todo figure out why shortID is used
			weightNew: u,
		})
	}

	// canonically sort validators
	// Note: validators are sorted by ID, sorting by weight would not create a
	// canonically sorted list
	sort.Sort(validators)

	// convert the slice of validators to a slice of weights
	validatorWeights := make([]uint64, len(validators))
	for i, v := range validators {
		validatorWeights[i] = v.weight
	}

	if err := w.sampler.Initialize(validatorWeights); err != nil {
		return 0, err
	}

	numToSample := MaxWindows
	if weight < uint64(numToSample) {
		numToSample = int(weight)
	}

	seed := chainHeight ^ w.chainSource
	w.sampler.Seed(int64(seed))

	indices, err := w.sampler.Sample(numToSample)
	if err != nil {
		return 0, err
	}

	delay := time.Duration(0)
	for _, index := range indices {
		nodeID := validators[index].id
		if nodeID == validatorID {
			return delay, nil
		}
		delay += WindowDuration
	}
	return delay, nil
}

func (w *windower) GetValidators(hash []byte) (map[string]float64, error) {
	var id [32]byte
	copy(id[:], hash)

	m, err := w.valClient.GetValidators(id)
	if err != nil {
		return nil, err
	}
	return convertShortIdmapToStringMap(m), nil

}

func convertShortIdmapToStringMap(m map[ids.ShortID]float64) map[string]float64 {
	retM := make(map[string]float64)
	for key, val := range m {
		retM[key.String()] = val
	}
	return retM
}
