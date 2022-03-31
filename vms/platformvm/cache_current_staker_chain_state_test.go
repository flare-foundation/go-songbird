// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
)

func TestPrimaryValidatorSet(t *testing.T) {
	// Initialize the chain state
	nodeID1 := ids.GenerateTestShortID()
	node0Weight := uint64(1)
	validator1 := &currentValidatorImpl{
		addValidatorTx: &UnsignedAddValidatorTx{
			Validator: Validator{
				Wght: node0Weight,
			},
		},
	}

	nodeID2 := ids.GenerateTestShortID()
	node1Weight := uint64(2)
	validator2 := &currentValidatorImpl{
		addValidatorTx: &UnsignedAddValidatorTx{
			Validator: Validator{
				Wght: node1Weight,
			},
		},
	}

	nodeID3 := ids.GenerateTestShortID()
	node2Weight := uint64(2)
	validator3 := &currentValidatorImpl{
		addValidatorTx: &UnsignedAddValidatorTx{
			Validator: Validator{
				Wght: node2Weight,
			},
		},
	}

	cs := &currentStakerChainStateImpl{
		validatorsByNodeID: map[ids.ShortID]*currentValidatorImpl{
			nodeID1: validator1,
			nodeID2: validator2,
			nodeID3: validator3,
		},
	}
	nodeID4 := ids.GenerateTestShortID()

	{
		// Apply the on-chain validator set to [validators]
		validators, err := cs.ValidatorSet(constants.PrimaryNetworkID)
		assert.NoError(t, err)

		// Validate that the state was applied and the old state was cleared
		assert.EqualValues(t, 3, validators.Len())
		assert.EqualValues(t, node0Weight+node1Weight+node2Weight, validators.Weight())
		gotNode0Weight, exists := validators.GetWeight(nodeID1)
		assert.True(t, exists)
		assert.EqualValues(t, node0Weight, gotNode0Weight)
		gotNode1Weight, exists := validators.GetWeight(nodeID2)
		assert.True(t, exists)
		assert.EqualValues(t, node1Weight, gotNode1Weight)
		gotNode2Weight, exists := validators.GetWeight(nodeID3)
		assert.True(t, exists)
		assert.EqualValues(t, node2Weight, gotNode2Weight)
		_, exists = validators.GetWeight(nodeID4)
		assert.False(t, exists)
	}

	{
		// Apply the on-chain validator set again
		validators, err := cs.ValidatorSet(constants.PrimaryNetworkID)
		assert.NoError(t, err)

		// The state should be the same
		assert.EqualValues(t, 3, validators.Len())
		assert.EqualValues(t, node0Weight+node1Weight+node2Weight, validators.Weight())
		gotNode0Weight, exists := validators.GetWeight(nodeID1)
		assert.True(t, exists)
		assert.EqualValues(t, node0Weight, gotNode0Weight)
		gotNode1Weight, exists := validators.GetWeight(nodeID2)
		assert.True(t, exists)
		assert.EqualValues(t, node1Weight, gotNode1Weight)
		gotNode2Weight, exists := validators.GetWeight(nodeID3)
		assert.True(t, exists)
		assert.EqualValues(t, node2Weight, gotNode2Weight)
	}
}

func TestSubnetValidatorSet(t *testing.T) {
	subnetID := ids.GenerateTestID()

	// Initialize the chain state
	nodeID1 := ids.GenerateTestShortID()
	node0Weight := uint64(1)
	validator1 := &currentValidatorImpl{
		validatorImpl: validatorImpl{
			subnets: map[ids.ID]*UnsignedAddSubnetValidatorTx{
				subnetID: {
					Validator: SubnetValidator{
						Validator: Validator{
							Wght: node0Weight,
						},
					},
				},
			},
		},
	}

	nodeID2 := ids.GenerateTestShortID()
	node1Weight := uint64(2)
	validator2 := &currentValidatorImpl{
		validatorImpl: validatorImpl{
			subnets: map[ids.ID]*UnsignedAddSubnetValidatorTx{
				subnetID: {
					Validator: SubnetValidator{
						Validator: Validator{
							Wght: node1Weight,
						},
					},
				},
			},
		},
	}

	nodeID3 := ids.GenerateTestShortID()
	node2Weight := uint64(2)
	validator3 := &currentValidatorImpl{
		validatorImpl: validatorImpl{
			subnets: map[ids.ID]*UnsignedAddSubnetValidatorTx{
				subnetID: {
					Validator: SubnetValidator{
						Validator: Validator{
							Wght: node2Weight,
						},
					},
				},
			},
		},
	}

	cs := &currentStakerChainStateImpl{
		validatorsByNodeID: map[ids.ShortID]*currentValidatorImpl{
			nodeID1: validator1,
			nodeID2: validator2,
			nodeID3: validator3,
		},
	}

	nodeID4 := ids.GenerateTestShortID()

	{
		// Apply the on-chain validator set to [validators]
		validators, err := cs.ValidatorSet(subnetID)
		assert.NoError(t, err)

		// Validate that the state was applied and the old state was cleared
		assert.EqualValues(t, 3, validators.Len())
		assert.EqualValues(t, node0Weight+node1Weight+node2Weight, validators.Weight())
		gotNode0Weight, exists := validators.GetWeight(nodeID1)
		assert.True(t, exists)
		assert.EqualValues(t, node0Weight, gotNode0Weight)
		gotNode1Weight, exists := validators.GetWeight(nodeID2)
		assert.True(t, exists)
		assert.EqualValues(t, node1Weight, gotNode1Weight)
		gotNode2Weight, exists := validators.GetWeight(nodeID3)
		assert.True(t, exists)
		assert.EqualValues(t, node2Weight, gotNode2Weight)
		_, exists = validators.GetWeight(nodeID4)
		assert.False(t, exists)
	}

	{
		// Apply the on-chain validator set again
		validators, err := cs.ValidatorSet(subnetID)
		assert.NoError(t, err)

		// The state should be the same
		assert.EqualValues(t, 3, validators.Len())
		assert.EqualValues(t, node0Weight+node1Weight+node2Weight, validators.Weight())
		gotNode0Weight, exists := validators.GetWeight(nodeID1)
		assert.True(t, exists)
		assert.EqualValues(t, node0Weight, gotNode0Weight)
		gotNode1Weight, exists := validators.GetWeight(nodeID2)
		assert.True(t, exists)
		assert.EqualValues(t, node1Weight, gotNode1Weight)
		gotNode2Weight, exists := validators.GetWeight(nodeID3)
		assert.True(t, exists)
		assert.EqualValues(t, node2Weight, gotNode2Weight)
	}
}
