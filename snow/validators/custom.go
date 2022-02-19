package validators

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
)

func custom() Set {
	weight := uint64(200000)
	path := os.Getenv("VALIDATORS")
	if path == "" {
		panic("custom validator file path not defined")
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("could not read custom validator data (path: %s): %s", path, err))
	}
	var nodeIDs []string
	err = json.Unmarshal(data, &nodeIDs)
	if err != nil {
		panic(fmt.Sprintf("could not decode custom validator datas: %s", err))
	}
	set := NewSet()
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
