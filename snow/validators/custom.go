package validators

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
)

type FlareSet struct {
	Validators []FlareValidator `json:"validators"`
}

type FlareValidator struct {
	NodeID string `json:"nodeID"`
	Weight uint64 `json:"weight"`
}

var custom Set

func init() {
	custom = NewSet()
	validatorPath := os.Getenv("VALIDATORS")
	if validatorPath == "" {
		return
	}
	validatorData, err := ioutil.ReadFile(validatorPath)
	if err != nil {
		panic(fmt.Sprintf("could not read quorum data (path: %s): %s", validatorPath, err))
	}
	var set FlareSet
	err = json.Unmarshal(validatorData, &set)
	if err != nil {
		panic(fmt.Sprintf("could not decode quorum: %s", err))
	}
	for _, validator := range set.Validators {
		nodeID, err := ids.ShortFromPrefixedString(validator.NodeID, constants.NodeIDPrefix)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = custom.AddWeight(nodeID, validator.Weight)
		if err != nil {
			panic(fmt.Sprintf("could not add weight for validator (node: %x, weight: %d): %s", nodeID, validator.Weight, err))
		}
	}
}
