// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:build integration
// +build integration

package chains

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow"
	"github.com/flare-foundation/flare/snow/validation"
	"github.com/flare-foundation/flare/utils/logging"
	"github.com/flare-foundation/flare/vms/rpcchainvm"
)

func TestManagerIntegration(t *testing.T) {
	cfg, _ := logging.DefaultConfig()
	log, _ := logging.NewTestLog(cfg)
	ctx := &snow.Context{
		Log: log,
	}

	vmFactory := &rpcchainvm.Factory{
		Path: filepath.Join("../build/plugins/mock"),
	}
	vm, err := vmFactory.New(ctx)
	require.NoErrorf(t, err, "if this error occurs during testing, it is most likely that you need to run the `build_plugin_mock.sh` script")

	m := &manager{
		ManagerConfig: ManagerConfig{
			Validators: validation.NewSet(),
		},
	}

	r, ok := vm.(validation.Retriever)
	require.True(t, ok)
	ctx.ValidatorsRetriever = validation.NewCachingRetriever(r)
	ctx.ValidatorsUpdater = validation.NewUpdater(m.Validators, ctx.ValidatorsRetriever)

	got, err := ctx.ValidatorsRetriever.GetValidators(ids.ID{})
	require.NoError(t, err)

	assert.Equal(t, got.Len(), 9)
	for i, v := range got.List() {
		id := v.ID()
		assert.Equal(t, byte(i)+1, id[0])
		assert.Equal(t, uint64(i)+1, v.Weight())
	}
}
