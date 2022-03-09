package validators

import (
	"errors"
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
	if T.GetValidatorsByBlockIDF != nil {
		return T.GetValidatorsByBlockIDF(blockID)
	}

	return nil, errGetValidatorsByBlockID
}
