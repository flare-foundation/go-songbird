// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"time"

	"github.com/flare-foundation/flare/api"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
	"github.com/flare-foundation/flare/utils/formatting"
	cjson "github.com/flare-foundation/flare/utils/json"
	"github.com/flare-foundation/flare/utils/rpc"
)

type Client struct {
	requester rpc.EndpointRequester
}

// NewClient returns a Client for interacting with the P Chain endpoint
func NewClient(uri string, requestTimeout time.Duration) Client {
	return &client{
		requester: rpc.NewEndpointRequester(uri, "/ext/P", "platform", requestTimeout),
	}
}

func (c *client) GetHeight() (uint64, error) {
	res := &GetHeightResponse{}
	err := c.requester.SendRequest("getHeight", struct{}{}, res)
	return uint64(res.Height), err
}

func (c *client) ExportKey(user api.UserPass, address string) (string, error) {
	res := &ExportKeyReply{}
	err := c.requester.SendRequest("exportKey", &ExportKeyArgs{
		UserPass: user,
		Address:  address,
	}, res)
	return res.PrivateKey, err
}

func (c *client) ImportKey(user api.UserPass, privateKey string) (string, error) {
	res := &api.JSONAddress{}
	err := c.requester.SendRequest("importKey", &ImportKeyArgs{
		UserPass:   user,
		PrivateKey: privateKey,
	}, res)
	return res.Address, err
}

func (c *client) GetBalance(address string) (*GetBalanceResponse, error) {
	res := &GetBalanceResponse{}
	err := c.requester.SendRequest("getBalance", &api.JSONAddress{
		Address: address,
	}, res)
	return res, err
}

func (c *client) CreateAddress(user api.UserPass) (string, error) {
	res := &api.JSONAddress{}
	err := c.requester.SendRequest("createAddress", &user, res)
	return res.Address, err
}

func (c *client) ListAddresses(user api.UserPass) ([]string, error) {
	res := &api.JSONAddresses{}
	err := c.requester.SendRequest("listAddresses", &user, res)
	return res.Addresses, err
}

func (c *client) GetUTXOs(addrs []string, limit uint32, startAddress, startUTXOID string) ([][]byte, api.Index, error) {
	return c.GetAtomicUTXOs(addrs, "", limit, startAddress, startUTXOID)
}

func (c *client) GetAtomicUTXOs(addrs []string, sourceChain string, limit uint32, startAddress, startUTXOID string) ([][]byte, api.Index, error) {
	res := &api.GetUTXOsReply{}
	err := c.requester.SendRequest("getUTXOs", &api.GetUTXOsArgs{
		Addresses:   addrs,
		SourceChain: sourceChain,
		Limit:       cjson.Uint32(limit),
		StartIndex: api.Index{
			Address: startAddress,
			UTXO:    startUTXOID,
		},
		Encoding: formatting.Hex,
	}, res)
	if err != nil {
		return nil, api.Index{}, err
	}

	utxos := make([][]byte, len(res.UTXOs))
	for i, utxo := range res.UTXOs {
		utxoBytes, err := formatting.Decode(res.Encoding, utxo)
		if err != nil {
			return nil, api.Index{}, err
		}
		utxos[i] = utxoBytes
	}
	return utxos, res.EndIndex, nil
}

func (c *client) GetSubnets(ids []ids.ID) ([]APISubnet, error) {
	res := &GetSubnetsResponse{}
	err := c.requester.SendRequest("getSubnets", &GetSubnetsArgs{
		IDs: ids,
	}, res)
	return res.Subnets, err
}

func (c *client) GetStakingAssetID(subnetID ids.ID) (ids.ID, error) {
	res := &GetStakingAssetIDResponse{}
	err := c.requester.SendRequest("getStakingAssetID", &GetStakingAssetIDArgs{
		SubnetID: subnetID,
	}, res)
	return res.AssetID, err
}

func (c *client) GetCurrentValidators(subnetID ids.ID, nodeIDs []ids.ShortID) ([]interface{}, error) {
	nodeIDsStr := []string{}
	for _, nodeID := range nodeIDs {
		nodeIDsStr = append(nodeIDsStr, nodeID.PrefixedString(constants.NodeIDPrefix))
	}
	res := &GetCurrentValidatorsReply{}
	err := c.requester.SendRequest("getCurrentValidators", &GetCurrentValidatorsArgs{
		SubnetID: subnetID,
		NodeIDs:  nodeIDsStr,
	}, res)
	return res.Validators, err
}

func (c *client) GetPendingValidators(subnetID ids.ID, nodeIDs []ids.ShortID) ([]interface{}, []interface{}, error) {
	nodeIDsStr := []string{}
	for _, nodeID := range nodeIDs {
		nodeIDsStr = append(nodeIDsStr, nodeID.PrefixedString(constants.NodeIDPrefix))
	}
	res := &GetPendingValidatorsReply{}
	err := c.requester.SendRequest("getPendingValidators", &GetPendingValidatorsArgs{
		SubnetID: subnetID,
		NodeIDs:  nodeIDsStr,
	}, res)
	return res.Validators, res.Delegators, err
}

func (c *client) GetCurrentSupply() (uint64, error) {
	res := &GetCurrentSupplyReply{}
	err := c.requester.SendRequest("getCurrentSupply", struct{}{}, res)
	return uint64(res.Supply), err
}

func (c *client) SampleValidators(subnetID ids.ID, sampleSize uint16) ([]string, error) {
	res := &SampleValidatorsReply{}
	err := c.requester.SendRequest("sampleValidators", &SampleValidatorsArgs{
		SubnetID: subnetID,
		Size:     cjson.Uint16(sampleSize),
	}, res)
	return res.Validators, err
}

func (c *client) AddValidator(
	user api.UserPass,
	from []string,
	changeAddr string,
	rewardAddress,
	nodeID string,
	stakeAmount,
	startTime,
	endTime uint64,
	delegationFeeRate float32,
) (ids.ID, error) {
	res := &api.JSONTxID{}
	jsonStakeAmount := cjson.Uint64(stakeAmount)
	err := c.requester.SendRequest("addValidator", &AddValidatorArgs{
		JSONSpendHeader: api.JSONSpendHeader{
			UserPass:      user,
			JSONFromAddrs: api.JSONFromAddrs{From: from},
		},
		APIStaker: APIStaker{
			NodeID:      nodeID,
			StakeAmount: &jsonStakeAmount,
			StartTime:   cjson.Uint64(startTime),
			EndTime:     cjson.Uint64(endTime),
		},
		RewardAddress:     rewardAddress,
		DelegationFeeRate: cjson.Float32(delegationFeeRate),
	}, res)
	return res.TxID, err
}

func (c *client) AddDelegator(
	user api.UserPass,
	from []string,
	changeAddr string,
	rewardAddress,
	nodeID string,
	stakeAmount,
	startTime,
	endTime uint64,
) (ids.ID, error) {
	res := &api.JSONTxID{}
	jsonStakeAmount := cjson.Uint64(stakeAmount)
	err := c.requester.SendRequest("addDelegator", &AddDelegatorArgs{
		JSONSpendHeader: api.JSONSpendHeader{
			UserPass:       user,
			JSONFromAddrs:  api.JSONFromAddrs{From: from},
			JSONChangeAddr: api.JSONChangeAddr{ChangeAddr: changeAddr},
		}, APIStaker: APIStaker{
			NodeID:      nodeID,
			StakeAmount: &jsonStakeAmount,
			StartTime:   cjson.Uint64(startTime),
			EndTime:     cjson.Uint64(endTime),
		},
		RewardAddress: rewardAddress,
	}, res)
	return res.TxID, err
}

func (c *client) AddSubnetValidator(
	user api.UserPass,
	from []string,
	changeAddr string,
	subnetID,
	nodeID string,
	stakeAmount,
	startTime,
	endTime uint64,
) (ids.ID, error) {
	res := &api.JSONTxID{}
	jsonStakeAmount := cjson.Uint64(stakeAmount)
	err := c.requester.SendRequest("addSubnetValidator", &AddSubnetValidatorArgs{
		JSONSpendHeader: api.JSONSpendHeader{
			UserPass:       user,
			JSONFromAddrs:  api.JSONFromAddrs{From: from},
			JSONChangeAddr: api.JSONChangeAddr{ChangeAddr: changeAddr},
		},
		APIStaker: APIStaker{
			NodeID:      nodeID,
			StakeAmount: &jsonStakeAmount,
			StartTime:   cjson.Uint64(startTime),
			EndTime:     cjson.Uint64(endTime),
		},
		SubnetID: subnetID,
	}, res)
	return res.TxID, err
}

func (c *client) CreateSubnet(
	user api.UserPass,
	from []string,
	changeAddr string,
	controlKeys []string,
	threshold uint32,
) (ids.ID, error) {
	res := &api.JSONTxID{}
	err := c.requester.SendRequest("createSubnet", &CreateSubnetArgs{
		JSONSpendHeader: api.JSONSpendHeader{
			UserPass:       user,
			JSONFromAddrs:  api.JSONFromAddrs{From: from},
			JSONChangeAddr: api.JSONChangeAddr{ChangeAddr: changeAddr},
		},
		APISubnet: APISubnet{
			ControlKeys: controlKeys,
			Threshold:   cjson.Uint32(threshold),
		},
	}, res)
	return res.TxID, err
}

func (c *client) ExportAVAX(
	user api.UserPass,
	from []string,
	changeAddr string,
	to string,
	amount uint64,
) (ids.ID, error) {
	res := &api.JSONTxID{}
	err := c.requester.SendRequest("exportAVAX", &ExportAVAXArgs{
		JSONSpendHeader: api.JSONSpendHeader{
			UserPass:       user,
			JSONFromAddrs:  api.JSONFromAddrs{From: from},
			JSONChangeAddr: api.JSONChangeAddr{ChangeAddr: changeAddr},
		},
		To:     to,
		Amount: cjson.Uint64(amount),
	}, res)
	return res.TxID, err
}

func (c *client) ImportAVAX(
	user api.UserPass,
	from []string,
	changeAddr,
	to,
	sourceChain string,
) (ids.ID, error) {
	res := &api.JSONTxID{}
	err := c.requester.SendRequest("importAVAX", &ImportAVAXArgs{
		JSONSpendHeader: api.JSONSpendHeader{
			UserPass:       user,
			JSONFromAddrs:  api.JSONFromAddrs{From: from},
			JSONChangeAddr: api.JSONChangeAddr{ChangeAddr: changeAddr},
		},
		To:          to,
		SourceChain: sourceChain,
	}, res)
	return res.TxID, err
}

func (c *client) CreateBlockchain(
	user api.UserPass,
	from []string,
	changeAddr string,
	subnetID ids.ID,
	vmID string,
	fxIDs []string,
	name string,
	genesisData []byte,
) (ids.ID, error) {
	genesisDataStr, err := formatting.EncodeWithChecksum(formatting.Hex, genesisData)
	if err != nil {
		return ids.ID{}, err
	}

	res := &api.JSONTxID{}
	err = c.requester.SendRequest("createBlockchain", &CreateBlockchainArgs{
		JSONSpendHeader: api.JSONSpendHeader{
			UserPass:       user,
			JSONFromAddrs:  api.JSONFromAddrs{From: from},
			JSONChangeAddr: api.JSONChangeAddr{ChangeAddr: changeAddr},
		},
		SubnetID:    subnetID,
		VMID:        vmID,
		FxIDs:       fxIDs,
		Name:        name,
		GenesisData: genesisDataStr,
		Encoding:    formatting.Hex,
	}, res)
	return res.TxID, err
}

func (c *client) GetBlockchainStatus(blockchainID string) (BlockchainStatus, error) {
	res := &GetBlockchainStatusReply{}
	err := c.requester.SendRequest("getBlockchainStatus", &GetBlockchainStatusArgs{
		BlockchainID: blockchainID,
	}, res)
	return res.Status, err
}

func (c *client) ValidatedBy(blockchainID ids.ID) (ids.ID, error) {
	res := &ValidatedByResponse{}
	err := c.requester.SendRequest("validatedBy", &ValidatedByArgs{
		BlockchainID: blockchainID,
	}, res)
	return res.SubnetID, err
}

func (c *client) Validates(subnetID ids.ID) ([]ids.ID, error) {
	res := &ValidatesResponse{}
	err := c.requester.SendRequest("validates", &ValidatesArgs{
		SubnetID: subnetID,
	}, res)
	return res.BlockchainIDs, err
}

func (c *client) GetBlockchains() ([]APIBlockchain, error) {
	res := &GetBlockchainsResponse{}
	err := c.requester.SendRequest("getBlockchains", struct{}{}, res)
	return res.Blockchains, err
}

func (c *client) IssueTx(txBytes []byte) (ids.ID, error) {
	txStr, err := formatting.EncodeWithChecksum(formatting.Hex, txBytes)
	if err != nil {
		return ids.ID{}, err
	}

	res := &api.JSONTxID{}
	err = c.requester.SendRequest("issueTx", &api.FormattedTx{
		Tx:       txStr,
		Encoding: formatting.Hex,
	}, res)
	return res.TxID, err
}

func (c *client) GetTx(txID ids.ID) ([]byte, error) {
	res := &api.FormattedTx{}
	err := c.requester.SendRequest("getTx", &api.GetTxArgs{
		TxID:     txID,
		Encoding: formatting.Hex,
	}, res)
	if err != nil {
		return nil, err
	}
	return formatting.Decode(res.Encoding, res.Tx)
}

func (c *client) GetTxStatus(txID ids.ID, includeReason bool) (*GetTxStatusResponse, error) {
	res := new(GetTxStatusResponse)
	err := c.requester.SendRequest("getTxStatus", &GetTxStatusArgs{
		TxID:          txID,
		IncludeReason: includeReason,
	}, res)
	return res, err
}

func (c *client) GetStake(addrs []string) (*GetStakeReply, error) {
	res := new(GetStakeReply)
	err := c.requester.SendRequest("getStake", &api.JSONAddresses{
		Addresses: addrs,
	}, res)
	return res, err
}

func (c *client) GetMinStake() (uint64, uint64, error) {
	res := new(GetMinStakeReply)
	err := c.requester.SendRequest("getMinStake", struct{}{}, res)
	return uint64(res.MinValidatorStake), uint64(res.MinDelegatorStake), err
}

func (c *client) GetTotalStake() (uint64, error) {
	res := new(GetTotalStakeReply)
	err := c.requester.SendRequest("getTotalStake", struct{}{}, res)
	return uint64(res.Stake), err
}

func (c *client) GetMaxStakeAmount(subnetID ids.ID, nodeID string, startTime, endTime uint64) (uint64, error) {
	res := new(GetMaxStakeAmountReply)
	err := c.requester.SendRequest("getMaxStakeAmount", &GetMaxStakeAmountArgs{
		SubnetID:  subnetID,
		NodeID:    nodeID,
		StartTime: cjson.Uint64(startTime),
		EndTime:   cjson.Uint64(endTime),
	}, res)
	return uint64(res.Amount), err
}

func (c *client) GetRewardUTXOs(args *api.GetTxArgs) ([][]byte, error) {
	res := &GetRewardUTXOsReply{}
	err := c.requester.SendRequest("getRewardUTXOs", args, res)
	if err != nil {
		return nil, err
	}
	utxos := make([][]byte, len(res.UTXOs))
	for i, utxoStr := range res.UTXOs {
		utxoBytes, err := formatting.Decode(res.Encoding, utxoStr)
		if err != nil {
			return nil, err
		}
		utxos[i] = utxoBytes
	}
	return utxos, err
}

func (c *client) GetTimestamp() (time.Time, error) {
	res := &GetTimestampReply{}
	err := c.requester.SendRequest("getTimestamp", struct{}{}, res)
	return res.Timestamp, err
}

func (c *client) GetValidatorsAt(subnetID ids.ID, height uint64) (map[string]uint64, error) {
	res := &GetValidatorsAtReply{}
	err := c.requester.SendRequest("getValidatorsAt", &GetValidatorsAtArgs{
		SubnetID: subnetID,
		Height:   cjson.Uint64(height),
	}, res)
	return res.Validators, err
}
