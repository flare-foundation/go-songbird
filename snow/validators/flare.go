package validators

import (
	"fmt"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
)

func flare() Set {
	weight := uint64(50000)
	nodeIDs := []string{}
	set := NewSet()
	for _, nodeID := range nodeIDs {
		shortID, err := ids.ShortFromPrefixedString(nodeID, constants.NodeIDPrefix)
		if err != nil {
			panic(fmt.Sprintf("invalid flare validator node ID: %s", nodeID))
		}
		err = set.AddWeight(shortID, weight)
		if err != nil {
			panic(fmt.Sprintf("could not add weight for validator (node: %s, weight: %d): %s", nodeID, weight, err))
		}
	}
	return set
}
