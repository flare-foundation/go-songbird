package validators

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
)

func TestCachingRetriever_GetValidators(t *testing.T) {
	testBlockID := ids.ID{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
	}

	t.Run("cache call count", func(t *testing.T) {
		t.Parallel()

		callsCount := 0
		cacheCallCount := 0
		validators := make(map[ids.ID]Set)
		retrieverMock := &TestRetriever{
			GetValidatorsByBlockIDFunc: func(blockID ids.ID) (Set, error) {
				callsCount++

				if set, ok := validators[blockID]; ok {
					cacheCallCount++
					return set, nil
				}
				validators[blockID] = loadCostonValidators()
				return validators[blockID], nil
			},
		}
		cr := NewCachingRetriever(retrieverMock)
		set, err := cr.GetValidators(testBlockID)
		assert.NoError(t, err)
		assert.Equal(t, loadCostonValidators(), set)
		assert.Equal(t, 1, callsCount)
		assert.Equal(t, 0, cacheCallCount)

		set, err = cr.GetValidators(testBlockID)
		assert.Equal(t, 1, callsCount)
		assert.Equal(t, 0, cacheCallCount)
		assert.Equal(t, uint64(1000000), set.Weight())

	})

	t.Run("error call", func(t *testing.T) {
		validators := make(map[ids.ID]Set)
		retrieverMock := &TestRetriever{
			GetValidatorsByBlockIDFunc: func(blockID ids.ID) (Set, error) {
				if blockID == testBlockID {
					return nil, fmt.Errorf("Couldn't get validators")
				}
				validators[blockID] = loadCostonValidators()
				return validators[blockID], nil
			},
		}
		cr := NewCachingRetriever(retrieverMock)
		_, err := cr.GetValidators(testBlockID)
		assert.Error(t, err)

	})
}

func loadCostonValidators() Set {
	weight := uint64(200_000)
	nodeIDs := []string{
		"NodeID-5dDZXn99LCkDoEi6t9gTitZuQmhokxQTc",
		"NodeID-EkH8wyEshzEQBToAdR7Fexxcj9rrmEEHZ",
		"NodeID-FPAwqHjs8Mw8Cuki5bkm3vSVisZr8t2Lu",
		"NodeID-AQghDJTU3zuQj73itPtfTZz6CxsTQVD3R",
		"NodeID-HaZ4HpanjndqSuN252chFsTysmdND5meA",
	}
	set := NewSet()
	for _, nodeID := range nodeIDs {
		shortID, err := ids.ShortFromPrefixedString(nodeID, constants.NodeIDPrefix)
		if err != nil {
			panic(fmt.Sprintf("invalid coston validator node ID: %s", nodeID))
		}
		err = set.AddWeight(shortID, weight)
		if err != nil {
			panic(fmt.Sprintf("could not add weight for validator (node: %s, weight: %d): %s", nodeID, weight, err))
		}
	}
	return set
}
