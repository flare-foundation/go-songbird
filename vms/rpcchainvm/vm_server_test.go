// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcchainvm

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare/api/proto/vmproto"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow/validation"
)

func TestVMServer_FetchValidators(t *testing.T) {
	testBlockID := ids.ID{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
	}

	// Generate fake validator set.
	validIDs := generateValidIDs(5)
	set := validation.NewSet()
	for i := range validIDs {
		id, _ := ids.ToShortID(validIDs[i])
		err := set.AddWeight(id, 5)
		require.NoError(t, err)
	}

	tests := []struct {
		name             string
		blockID          []byte
		set              validation.Set
		retrieverFailure bool
		wantErr          require.ErrorAssertionFunc
	}{
		{
			name:             "nominal case",
			blockID:          testBlockID[:],
			set:              set,
			retrieverFailure: false,
			wantErr:          require.NoError,
		},
		{
			name:             "empty validator list",
			blockID:          testBlockID[:],
			set:              validation.NewSet(),
			retrieverFailure: false,
			wantErr:          require.NoError,
		},
		{
			name:             "invalid block ID",
			blockID:          []byte{0x01, 0x02},
			set:              set,
			retrieverFailure: false,
			wantErr:          require.Error,
		},
		{
			name:             "validator retriever failure",
			blockID:          testBlockID[:],
			set:              set,
			retrieverFailure: true,
			wantErr:          require.Error,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			// Set up mock to assert that the expected block ID is given, and
			// either return an error if test.retrieverFailure is set to true,
			// or return the test set otherwise.
			chainVMMock := ChainVMMock{
				GetValidatorsFunc: func(gotID ids.ID) (validation.Set, error) {
					assert.Equal(t, testBlockID, gotID)

					if test.retrieverFailure {
						return nil, errors.New("dummy error")
					}

					return test.set, nil
				},
			}
			// Create a VMServer using the mock.
			server := NewServer(chainVMMock, nil)

			resp, err := server.FetchValidators(context.Background(), &vmproto.FetchValidatorsRequest{
				BlkId: test.blockID[:],
			})
			test.wantErr(t, err)

			// If no error, assert that the response matches with the expected
			// output for this test.
			if err == nil {
				assert.Len(t, resp.Weights, test.set.Len())
				assert.Len(t, resp.ValidatorIds, test.set.Len())

				want := test.set.List()
				for i := range want {
					wantID := want[i].ID()
					assert.Equal(t, wantID[:], resp.ValidatorIds[i][:])
					assert.Equal(t, want[i].Weight(), resp.Weights[i])
				}
			}
		})
	}
}
