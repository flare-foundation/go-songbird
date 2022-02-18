package validators

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
)

type FlareSet struct {
	Validators []FlareValidator `json:"validators"`
}

type FlareValidator struct {
	NodeID string `json:"nodeID"`
	Weight uint64 `json:"weight"`
}

var FBA Set

func init() {
	FBA = NewSet()
	quorumPath := os.Getenv("FBA_VALs")
	if quorumPath == "" {
		return
	}
	quorumData, err := ioutil.ReadFile(quorumPath)
	if err != nil {
		panic(fmt.Sprintf("could not read quorum data (path: %s): %s", quorumPath, err))
	}
	var set FlareSet
	err = json.Unmarshal(quorumData, &set)
	if err != nil {
		panic(fmt.Sprintf("could not decode quorum: %s", err))
	}
	for _, validator := range set.Validators {
		nodeID, err := ids.ShortFromPrefixedString(validator.NodeID, constants.NodeIDPrefix)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = FBA.AddWeight(nodeID, validator.Weight)
		if err != nil {
			panic(fmt.Sprintf("could not add weight for validator (node: %x, weight: %d): %s", nodeID, validator.Weight, err))
		}
	}
}
