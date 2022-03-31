// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"errors"
	"fmt"
	"time"

	"github.com/gorilla/rpc/v2"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/flare-foundation/flare/cache"
	"github.com/flare-foundation/flare/chains"
	"github.com/flare-foundation/flare/codec"
	"github.com/flare-foundation/flare/codec/linearcodec"
	"github.com/flare-foundation/flare/database/manager"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow"
	"github.com/flare-foundation/flare/snow/choices"
	"github.com/flare-foundation/flare/snow/consensus/snowman"
	"github.com/flare-foundation/flare/snow/engine/common"
	"github.com/flare-foundation/flare/snow/engine/snowman/block"
	"github.com/flare-foundation/flare/snow/uptime"
	"github.com/flare-foundation/flare/snow/validation"
	"github.com/flare-foundation/flare/utils"
	"github.com/flare-foundation/flare/utils/constants"
	"github.com/flare-foundation/flare/utils/crypto"
	"github.com/flare-foundation/flare/utils/json"
	"github.com/flare-foundation/flare/utils/logging"
	"github.com/flare-foundation/flare/utils/timer/mockable"
	"github.com/flare-foundation/flare/utils/wrappers"
	"github.com/flare-foundation/flare/version"
	"github.com/flare-foundation/flare/vms/components/avax"
	"github.com/flare-foundation/flare/vms/platformvm/reward"
	"github.com/flare-foundation/flare/vms/secp256k1fx"

	safemath "github.com/flare-foundation/flare/utils/math"
)

const (
	droppedTxCacheSize = 64

	// MaxValidatorWeightFactor is the maximum factor of the validator stake
	// that is allowed to be placed on a validator.
	MaxValidatorWeightFactor uint64 = 5

	// Maximum future start time for staking/delegating
	maxFutureStartTime = 24 * 7 * 2 * time.Hour
)

var (
	errInvalidID         = errors.New("invalid ID")
	errDSCantValidate    = errors.New("new blockchain can't be validated by primary network")
	errStartTimeTooEarly = errors.New("start time is before the current chain time")
	errStartAfterEndTime = errors.New("start time is after the end time")
	errWrongCacheType    = errors.New("unexpectedly cached type")

	_ block.ChainVM        = &VM{}
	_ validation.Connector = &VM{}
	_ secp256k1fx.VM       = &VM{}
	_ Fx                   = &secp256k1fx.Fx{}
)

type VM struct {
	Factory
	metrics
	avax.AddressManager
	avax.AtomicUTXOManager
	*network

	// Used to get time. Useful for faking time during tests.
	clock mockable.Clock

	// Used to create and use keys.
	factory crypto.FactorySECP256K1R

	blockBuilder blockBuilder

	uptimeManager uptime.Manager

	rewards reward.Calculator

	// The context of this vm
	ctx       *snow.Context
	dbManager manager.Manager

	internalState InternalState

	// ID of the preferred block
	preferred ids.ID

	// ID of the last accepted block
	lastAcceptedID ids.ID

	fx            Fx
	codecRegistry codec.Registry

	// Bootstrapped remembers if this chain has finished bootstrapping or not
	bootstrapped utils.AtomicBool

	// Contains the IDs of transactions recently dropped because they failed
	// verification. These txs may be re-issued and put into accepted blocks, so
	// check the database to see if it was later committed/aborted before
	// reporting that it's dropped.
	// Key: Tx ID
	// Value: String repr. of the verification error
	droppedTxCache cache.LRU

	// Maps caches for each subnet that is currently whitelisted.
	// Key: Subnet ID
	// Value: cache mapping height -> validator set map
	validatorSetCaches map[ids.ID]cache.Cacher

	// Key: block ID
	// Value: the block
	currentBlocks map[ids.ID]Block
}

// Initialize this blockchain.
// [vm.ChainManager] and [vm.vdrMgr] must be set before this function is called.
func (vm *VM) Initialize(
	ctx *snow.Context,
	dbManager manager.Manager,
	genesisBytes []byte,
	upgradeBytes []byte,
	configBytes []byte,
	toEngine chan<- common.Message,
	_ []*common.Fx,
	appSender common.AppSender,
) error {
	ctx.Log.Verbo("initializing platform chain")

	registerer := prometheus.NewRegistry()
	if err := ctx.Metrics.Register(registerer); err != nil {
		return err
	}

	// Initialize metrics as soon as possible
	if err := vm.metrics.Initialize("", registerer); err != nil {
		return err
	}

	// Initialize the utility to parse addresses
	vm.AddressManager = avax.NewAddressManager(ctx)

	// Initialize the utility to fetch atomic UTXOs
	vm.AtomicUTXOManager = avax.NewAtomicUTXOManager(ctx.SharedMemory, Codec)

	vm.fx = &secp256k1fx.Fx{}

	vm.ctx = ctx
	vm.dbManager = dbManager

	vm.codecRegistry = linearcodec.NewDefault()
	if err := vm.fx.Initialize(vm); err != nil {
		return err
	}

	vm.droppedTxCache = cache.LRU{Size: droppedTxCacheSize}
	vm.validatorSetCaches = make(map[ids.ID]cache.Cacher)
	vm.currentBlocks = make(map[ids.ID]Block)

	if err := vm.blockBuilder.Initialize(vm, toEngine, registerer); err != nil {
		return fmt.Errorf(
			"failed to initialize the block builder: %w",
			err,
		)
	}
	vm.network = newNetwork(vm.ApricotPhase4Time, appSender, vm)
	vm.rewards = reward.NewCalculator(vm.RewardConfig)

	is, err := NewMeteredInternalState(vm, vm.dbManager.Current().Database, genesisBytes, registerer)
	if err != nil {
		return err
	}
	vm.internalState = is

	// Initialize the utility to track validator uptimes
	vm.uptimeManager = uptime.NewManager(is)
	vm.UptimeLockedCalculator.SetCalculator(&vm.bootstrapped, &ctx.Lock, vm.uptimeManager)

	// Create all of the chains that the database says exist
	if err := vm.initBlockchains(); err != nil {
		return fmt.Errorf(
			"failed to initialize blockchains: %w",
			err,
		)
	}

	vm.lastAcceptedID = is.GetLastAccepted()

	ctx.Log.Info("initializing last accepted block as %s", vm.lastAcceptedID)

	// Build off the most recently accepted block
	return vm.SetPreference(vm.lastAcceptedID)
}

// Create all chains that exist that this node validates.
func (vm *VM) initBlockchains() error {
	if err := vm.createSubnet(constants.PrimaryNetworkID); err != nil {
		return err
	}

	if vm.StakingEnabled {
		for subnetID := range vm.WhitelistedSubnets {
			if err := vm.createSubnet(subnetID); err != nil {
				return err
			}
		}
	} else {
		subnets, err := vm.internalState.GetSubnets()
		if err != nil {
			return err
		}
		for _, subnet := range subnets {
			if err := vm.createSubnet(subnet.ID()); err != nil {
				return err
			}
		}
	}
	return nil
}

// Create the subnet with ID [subnetID]
func (vm *VM) createSubnet(subnetID ids.ID) error {
	chains, err := vm.internalState.GetChains(subnetID)
	if err != nil {
		return err
	}
	for _, chain := range chains {
		if err := vm.createChain(chain); err != nil {
			return err
		}
	}
	return nil
}

// Create the blockchain described in [tx], but only if this node is a member of
// the subnet that validates the chain
func (vm *VM) createChain(tx *Tx) error {
	unsignedTx, ok := tx.UnsignedTx.(*UnsignedCreateChainTx)
	if !ok {
		return errWrongTxType
	}

	if vm.StakingEnabled && // Staking is enabled, so nodes might not validate all chains
		constants.PrimaryNetworkID != unsignedTx.SubnetID && // All nodes must validate the primary network
		!vm.WhitelistedSubnets.Contains(unsignedTx.SubnetID) { // This node doesn't validate this blockchain
		return nil
	}

	chainParams := chains.ChainParameters{
		ID:          tx.ID(),
		SubnetID:    unsignedTx.SubnetID,
		GenesisData: unsignedTx.GenesisData,
		VMAlias:     unsignedTx.VMID.String(),
	}
	for _, fxID := range unsignedTx.FxIDs {
		chainParams.FxAliases = append(chainParams.FxAliases, fxID.String())
	}
	vm.Chains.CreateChain(chainParams)
	return nil
}

// onBootstrapStarted marks this VM as bootstrapping
func (vm *VM) onBootstrapStarted() error {
	vm.bootstrapped.SetValue(false)
	return vm.fx.Bootstrapping()
}

// onNormalOperationsStarted marks this VM as bootstrapped
func (vm *VM) onNormalOperationsStarted() error {
	if vm.bootstrapped.GetValue() {
		return nil
	}
	vm.bootstrapped.SetValue(true)

	if err := vm.fx.Bootstrapped(); err != nil {
		return err
	}
	return vm.internalState.Commit()
}

func (vm *VM) SetState(state snow.State) error {
	switch state {
	case snow.Bootstrapping:
		return vm.onBootstrapStarted()
	case snow.NormalOp:
		return vm.onNormalOperationsStarted()
	default:
		return snow.ErrUnknownState
	}
}

// Shutdown this blockchain
func (vm *VM) Shutdown() error {
	if vm.dbManager == nil {
		return nil
	}

	vm.blockBuilder.Shutdown()

	if vm.bootstrapped.GetValue() {
		validatorList := vm.Validators.List()

		validatorIDs := make([]ids.ShortID, len(validatorList))
		for i, vdr := range validatorList {
			validatorIDs[i] = vdr.ID()
		}

		if err := vm.uptimeManager.Shutdown(validatorIDs); err != nil {
			return err
		}
		if err := vm.internalState.Commit(); err != nil {
			return err
		}
	}

	errs := wrappers.Errs{}
	errs.Add(
		vm.internalState.Close(),
		vm.dbManager.Close(),
	)
	return errs.Err
}

// BuildBlock builds a block to be added to consensus
func (vm *VM) BuildBlock() (snowman.Block, error) { return vm.blockBuilder.BuildBlock() }

func (vm *VM) ParseBlock(b []byte) (snowman.Block, error) {
	var blk Block
	if _, err := Codec.Unmarshal(b, &blk); err != nil {
		return nil, err
	}
	if err := blk.initialize(vm, b, choices.Processing, blk); err != nil {
		return nil, err
	}

	// TODO: remove this to make ParseBlock stateless
	if block, err := vm.GetBlock(blk.ID()); err == nil {
		// If we have seen this block before, return it with the most up-to-date
		// info
		return block, nil
	}
	return blk, nil
}

func (vm *VM) GetBlock(blkID ids.ID) (snowman.Block, error) { return vm.getBlock(blkID) }

func (vm *VM) getBlock(blkID ids.ID) (Block, error) {
	// If block is in memory, return it.
	if blk, exists := vm.currentBlocks[blkID]; exists {
		return blk, nil
	}
	return vm.internalState.GetBlock(blkID)
}

// LastAccepted returns the block most recently accepted
func (vm *VM) LastAccepted() (ids.ID, error) {
	return vm.lastAcceptedID, nil
}

// SetPreference sets the preferred block to be the one with ID [blkID]
func (vm *VM) SetPreference(blkID ids.ID) error {
	if blkID == vm.preferred {
		// If the preference didn't change, then this is a noop
		return nil
	}
	vm.preferred = blkID
	vm.blockBuilder.ResetTimer()
	return nil
}

func (vm *VM) Preferred() (Block, error) {
	return vm.getBlock(vm.preferred)
}

func (vm *VM) Version() (string, error) {
	return version.Current.String(), nil
}

// CreateHandlers returns a map where:
// * keys are API endpoint extensions
// * values are API handlers
func (vm *VM) CreateHandlers() (map[string]*common.HTTPHandler, error) {
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	server.RegisterInterceptFunc(vm.metrics.apiRequestMetrics.InterceptRequest)
	server.RegisterAfterFunc(vm.metrics.apiRequestMetrics.AfterRequest)
	if err := server.RegisterService(&Service{vm: vm}, "platform"); err != nil {
		return nil, err
	}

	return map[string]*common.HTTPHandler{
		"": {
			Handler: server,
		},
	}, nil
}

// CreateStaticHandlers returns a map where:
// * keys are API endpoint extensions
// * values are API handlers
func (vm *VM) CreateStaticHandlers() (map[string]*common.HTTPHandler, error) {
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	if err := server.RegisterService(&StaticService{}, "platform"); err != nil {
		return nil, err
	}

	return map[string]*common.HTTPHandler{
		"": {
			LockOptions: common.NoLock,
			Handler:     server,
		},
	}, nil
}

func (vm *VM) Connected(vdrID ids.ShortID, _ version.Application) error {
	return vm.uptimeManager.Connect(vdrID)
}

func (vm *VM) Disconnected(vdrID ids.ShortID) error {
	if err := vm.uptimeManager.Disconnect(vdrID); err != nil {
		return err
	}
	return vm.internalState.Commit()
}

// GetCurrentHeight returns the height of the last accepted block
func (vm *VM) GetCurrentHeight() (uint64, error) {
	lastAccepted, err := vm.getBlock(vm.lastAcceptedID)
	if err != nil {
		return 0, err
	}
	return lastAccepted.Height(), nil
}

func (vm *VM) CodecRegistry() codec.Registry { return vm.codecRegistry }

func (vm *VM) Clock() *mockable.Clock { return &vm.clock }

func (vm *VM) Logger() logging.Logger { return vm.ctx.Log }

// Returns the percentage of the total stake on the Primary Network of nodes
// connected to this node.
func (vm *VM) getPercentConnected() (float64, error) {

	validatorList := vm.Validators.List()

	var (
		connectedStake uint64
		err            error
	)
	for _, vdr := range validatorList {
		if !vm.uptimeManager.IsConnected(vdr.ID()) {
			continue // not connected to us --> don't include
		}
		connectedStake, err = safemath.Add64(connectedStake, vdr.Weight())
		if err != nil {
			return 0, err
		}
	}
	return float64(connectedStake) / float64(vm.Validators.Weight()), nil
}
