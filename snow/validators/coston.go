package validators

import (
	"fmt"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
)

var coston Set

func init() {
	weight := uint64(200000)
	nodeIDs := []string{
		"NodeID-5dDZXn99LCkDoEi6t9gTitZuQmhokxQTc",
		"NodeID-EkH8wyEshzEQBToAdR7Fexxcj9rrmEEHZ",
		"NodeID-FPAwqHjs8Mw8Cuki5bkm3vSVisZr8t2Lu",
		"NodeID-AQghDJTU3zuQj73itPtfTZz6CxsTQVD3R",
		"NodeID-HaZ4HpanjndqSuN252chFsTysmdND5meA",
	}
	coston = NewSet()
	for _, nodeID := range nodeIDs {
		nodeID, err := ids.ShortFromPrefixedString(nodeID, constants.NodeIDPrefix)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = songbird.AddWeight(nodeID, weight)
		if err != nil {
			panic(fmt.Sprintf("could not add weight for validator (node: %x, weight: %d): %s", nodeID, weight, err))
		}
	}
}
