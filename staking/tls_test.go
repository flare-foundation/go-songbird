// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package staking

import (
	"crypto"
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ava-labs/avalanchego/utils/hashing"
)

func TestMakeKeys(t *testing.T) {
	assert := assert.New(t)

	cert, err := NewTLSCert()
	assert.NoError(err)

	msg := []byte(fmt.Sprintf("msg %d", time.Now().Unix()))
	msgHash := hashing.ComputeHash256(msg)

	sig, err := cert.PrivateKey.(crypto.Signer).Sign(rand.Reader, msgHash, crypto.SHA256)
	assert.NoError(err)

	err = cert.Leaf.CheckSignature(cert.Leaf.SignatureAlgorithm, msg, sig)
	assert.NoError(err)
}
