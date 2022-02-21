package validators

import (
	"fmt"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
)

const (
	flareValidatorWeight = 50_000
)

func loadFlareValidators() Set {
	weight := uint64(flareValidatorWeight)
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
