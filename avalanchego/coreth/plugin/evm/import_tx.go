// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"math/big"

	"github.com/flare-foundation/flare/coreth/core/state"
	"github.com/flare-foundation/flare/coreth/params"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/flare/chains/atomic"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow"
	"github.com/flare-foundation/flare/utils/crypto"
	"github.com/flare-foundation/flare/vms/components/avax"
	"github.com/flare-foundation/flare/vms/secp256k1fx"
)

// UnsignedImportTx is an unsigned ImportTx
type UnsignedImportTx struct {
	avax.Metadata
	// ID of the network on which this tx was issued
	NetworkID uint32 `serialize:"true" json:"networkID"`
	// ID of this blockchain.
	BlockchainID ids.ID `serialize:"true" json:"blockchainID"`
	// Which chain to consume the funds from
	SourceChain ids.ID `serialize:"true" json:"sourceChain"`
	// Inputs that consume UTXOs produced on the chain
	ImportedInputs []*avax.TransferableInput `serialize:"true" json:"importedInputs"`
	// Outputs
	Outs []EVMOutput `serialize:"true" json:"outputs"`
}

// InputUTXOs returns the UTXOIDs of the imported funds
func (tx *UnsignedImportTx) InputUTXOs() ids.Set {
	return ids.Set{}
}

// Verify this transaction is well-formed
func (tx *UnsignedImportTx) Verify(
	ctx *snow.Context,
	rules params.Rules,
) error {
	return errImportTxsDisabled
}

func (tx *UnsignedImportTx) GasUsed(fixedFee bool) (uint64, error) {
	return 0, errImportTxsDisabled
}

// Amount of [assetID] burned by this transaction
func (tx *UnsignedImportTx) Burned(assetID ids.ID) (uint64, error) {
	return 0, errImportTxsDisabled
}

// SemanticVerify this transaction is valid.
func (tx *UnsignedImportTx) SemanticVerify(
	vm *VM,
	stx *Tx,
	parent *Block,
	baseFee *big.Int,
	rules params.Rules,
) error {
	return errImportTxsDisabled
}

// Accept this transaction and spend imported inputs
// We spend imported UTXOs here rather than in semanticVerify because
// we don't want to remove an imported UTXO in semanticVerify
// only to have the transaction not be Accepted. This would be inconsistent.
// Recall that imported UTXOs are not kept in a versionDB.
func (tx *UnsignedImportTx) Accept() (ids.ID, *atomic.Requests, error) {
	return ids.ID{}, nil, errImportTxsDisabled
}

// newImportTx returns a new ImportTx
func (vm *VM) newImportTx(
	chainID ids.ID, // chain to import from
	to common.Address, // Address of recipient
	baseFee *big.Int, // fee to use post-AP3
	keys []*crypto.PrivateKeySECP256K1R, // Keys to import the funds
) (*Tx, error) {
	return nil, errImportTxsDisabled
}

// newImportTx returns a new ImportTx
func (vm *VM) newImportTxWithUTXOs(
	chainID ids.ID, // chain to import from
	to common.Address, // Address of recipient
	baseFee *big.Int, // fee to use post-AP3
	kc *secp256k1fx.Keychain, // Keychain to use for signing the atomic UTXOs
	atomicUTXOs []*avax.UTXO, // UTXOs to spend
) (*Tx, error) {
	return nil, errImportTxsDisabled
}

// EVMStateTransfer performs the state transfer to increase the balances of
// accounts accordingly with the imported EVMOutputs
func (tx *UnsignedImportTx) EVMStateTransfer(ctx *snow.Context, state *state.StateDB) error {
	return errImportTxsDisabled
}
