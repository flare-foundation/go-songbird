package validators

import (
	"fmt"
	"os"
	"strings"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
)

const (
	customValidatorWeight = 200_000
	customValidatorEnv    = "CUSTOM_VALIDATORS"
)

func loadCustomValidators() Set {
	set := NewSet()
	weight := uint64(customValidatorWeight)
	customValidatorList := os.Getenv(customValidatorEnv)
	if customValidatorList == "" {
		panic("environment variable for custom validators empty")
	}
	nodeIDs := strings.Split(customValidatorList, ",")
	for _, nodeID := range nodeIDs {
		shortID, err := ids.ShortFromPrefixedString(nodeID, constants.NodeIDPrefix)
		if err != nil {
			panic(fmt.Sprintf("invalid custom validator node ID: %s", nodeID))
		}
		err = set.AddWeight(shortID, weight)
		if err != nil {
			panic(fmt.Sprintf("could not add weight for validator (node: %x, weight: %d): %s", shortID, weight, err))
		}
	}
	return set
}
