// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcchainvm

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/flare-foundation/flare/api/proto/vmproto"
	"github.com/flare-foundation/flare/ids"
)

func TestVMClient_GetValidators(t *testing.T) {
	testBlockID := ids.ID{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
	}
	invalidID := []byte{0x00, 0x01}
	validIDs := generateValidIDs(5)

	tests := []struct {
		name         string
		validatorIDs [][]byte
		weights      []uint64
		fetchFailure bool
		wantErr      require.ErrorAssertionFunc
	}{
		{
			name:         "nominal case",
			validatorIDs: validIDs,
			weights:      []uint64{64, 32, 128, 192, 48},
			fetchFailure: false,
			wantErr:      require.NoError,
		},
		{
			name:         "no validators",
			validatorIDs: [][]byte{},
			weights:      []uint64{},
			fetchFailure: false,
			wantErr:      require.NoError,
		},
		{
			name:         "validator IDs and weights do not match",
			validatorIDs: [][]byte{validIDs[0]},
			weights:      []uint64{64, 32, 128, 192, 48},
			fetchFailure: false,
			wantErr:      require.Error,
		},
		{
			name:         "invalid ID(s) — Len != 20",
			validatorIDs: [][]byte{validIDs[0], invalidID, validIDs[2], invalidID, invalidID},
			weights:      []uint64{64, 32, 128, 192, 48},
			fetchFailure: false,
			wantErr:      require.Error,
		},
		{
			name:         "vm fails to fetch validators",
			validatorIDs: validIDs,
			weights:      []uint64{64, 32, 128, 192, 48},
			fetchFailure: true,
			wantErr:      require.Error,
		},
	}

	for _, test := range tests {

		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			// Set up mock to assert that the expected block ID is given, and
			// either return an error if test.fetchFailure is set to true,
			// or return a response with the test parameters otherwise.
			clientMock := VMClientMock{
				FetchValidatorsFunc: func(_ context.Context, req *vmproto.FetchValidatorsRequest, _ ...grpc.CallOption) (*vmproto.FetchValidatorsResponse, error) {
					assert.Equal(t, testBlockID[:], req.BlkId[:])

					if test.fetchFailure {
						return nil, errors.New("dummy error")
					}

					resp := vmproto.FetchValidatorsResponse{
						ValidatorIds: test.validatorIDs,
						Weights:      test.weights,
					}

					return &resp, nil
				},
			}
			// Create a VMClient using the mock.
			client := NewClient(clientMock, nil)

			set, err := client.GetValidators(testBlockID)
			test.wantErr(t, err)

			// If no error, assert that the returned set matches with the expected
			// output for this test.
			if err == nil {
				got := set.List()
				assert.Len(t, got, len(test.validatorIDs))
				for i := range got {
					id := got[i].ID()
					assert.Equal(t, test.validatorIDs[i][:], id[:])
					assert.Equal(t, test.weights[i], got[i].Weight())
				}
			}
		})
	}
}

// generateValidIDs generates n byte slices of size 20.
// These are intended to be compatible with the ids.ShortID type.
func generateValidIDs(n int) [][]byte {
	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		res[i] = make([]byte, 20)
		for j := 0; j < 20; j++ {
			res[i][j] = byte(i*20+j)
		}
	}
	return res
}
