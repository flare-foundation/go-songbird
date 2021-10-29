package platformvm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/validators"
	"github.com/flare-foundation/flare/utils/constants"
)

var fbaValidators validators.Set

func init() {
	fbaValidators = validators.NewSet()
	quorumPath := os.Getenv("FBA_VALs")
	if quorumPath == "" {
		return
	}
	quorumData, err := ioutil.ReadFile(quorumPath)
	if err != nil {
		panic(fmt.Sprintf("could not read quorum data (path: %s): %s", quorumPath, err))
	}
	var quorum validators.Quorum
	err = json.Unmarshal(quorumData, &quorum)
	if err != nil {
		panic(fmt.Sprintf("could not decode quorum: %s", err))
	}
	for _, validator := range quorum.Validators {
		nodeID, err := ids.ShortFromPrefixedString(validator.NodeID, constants.NodeIDPrefix)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = fbaValidators.AddWeight(nodeID, validator.Weight)
		if err != nil {
			panic(fmt.Sprintf("could not add weight for validator (node: %x, weight: %d): %s", nodeID, validator.Weight, err))
		}
	}
}
