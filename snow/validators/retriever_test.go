package validators

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
)

type testRetriever struct {
	m map[ids.ID]Set
}

var Counter = 0
var GetValidatorsCallCounter = 0

func NewTestRetriever() *testRetriever {
	return &testRetriever{
		m: make(map[ids.ID]Set),
	}
}

func (c *testRetriever) GetValidators(blockID ids.ID) (Set, error) {
	GetValidatorsCallCounter++
	Counter = 0
	if s, ok := c.m[blockID]; ok { //todo have a counter here
		return s, nil
	}
	Counter++
	c.m[blockID] = loadCostonValidators()
	return c.m[blockID], nil
}

func TestCachingRetriever_GetValidators(t *testing.T) {
	GetValidatorsCallCounter = 0
	testBlockID := ids.ID{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
	}

	//testBlockIDNonExistent := ids.ID{
	//	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
	//	0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
	//	0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
	//	0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1E,
	//}

	cr := NewCachingRetriever(NewTestRetriever())
	set, err := cr.GetValidators(testBlockID)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1000000), set.Weight())
	assert.Equal(t, Counter, 1)
	assert.Equal(t, GetValidatorsCallCounter, 1)

	set, err = cr.GetValidators(testBlockID)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1000000), set.Weight())
	//assert.Equal(t, Counter, 0)
	assert.Equal(t, GetValidatorsCallCounter, 1)

	//set, err = cr.GetValidators(testBlockIDNonExistent)
	//assert.NoError(t, err)
	//assert.Equal(t, NewSet(), set)
	//assert.Equal(t, Counter, 0)
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
