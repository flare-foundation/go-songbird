package validation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare/ids"
)

func TestNewRetrievingUpdater(t *testing.T) {
	testBlockID := ids.ID{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
	}
	costonSet := loadCostonValidators(t)

	t.Run("nominal case", func(t *testing.T) {
		t.Parallel()

		callsCount := 0

		retrieverMock := &TestRetriever{
			GetValidatorsByBlockIDFunc: func(blockID ids.ID) (Set, error) {
				callsCount++
				return costonSet, nil
			},
		}
		cr := NewCachingRetriever(retrieverMock)
		u := NewRetrievingUpdater(loadCostonValidators(t), cr)

		err := u.UpdateValidators(testBlockID)
		require.NoError(t, err)
		assert.Equal(t, 1, callsCount)

		set, err := cr.GetValidators(testBlockID)
		require.NoError(t, err)
		assert.Equal(t, costonSet, set)
		assert.Equal(t, 1, callsCount)
	})

	t.Run("handles failure to get validators", func(t *testing.T) {
		t.Parallel()

		retrieverMock := &TestRetriever{
			GetValidatorsByBlockIDFunc: func(blockID ids.ID) (Set, error) {
				return nil, fmt.Errorf("Couldn't get validators")
			},
		}
		cr := NewCachingRetriever(retrieverMock)
		u := NewRetrievingUpdater(NewSet(), cr)

		err := u.UpdateValidators(testBlockID)
		require.Error(t, err)
	})
}
