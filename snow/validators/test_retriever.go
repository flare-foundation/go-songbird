package validators

import (
	"errors"
	"fmt"
	"testing"

	"github.com/flare-foundation/flare/ids"
)

var (
	errGetValidatorsByBlockID = errors.New("unexpectedly called GetValidatorsByBlockID")
)

type TestRetriever struct {
	T *testing.T

	GetValidatorsByBlockIDF func(blockID ids.ID) (Set, error)
}

func (T *TestRetriever) GetValidatorsByBlockID(blockID ids.ID) (Set, error) {
	fmt.Println("TestRetriever GetValidatorsByBlockID called")
	if T.GetValidatorsByBlockIDF != nil {
		return T.GetValidatorsByBlockIDF(blockID)
	}
	//s := NewSet()
	//s.AddWeight(ids.ShortID{11}, 2)
	//return s, nil
	return nil, errGetValidatorsByBlockID
}
