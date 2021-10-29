// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package keystore

import (
	"github.com/flare-foundation/flare/database"
	"github.com/flare-foundation/flare/database/encdb"
	"github.com/flare-foundation/flare/ids"
)

var _ BlockchainKeystore = &blockchainKeystore{}

type BlockchainKeystore interface {
	// Get a database that is able to read and write unencrypted values from the
	// underlying database.
	GetDatabase(username, password string) (*encdb.Database, error)

	// Get the underlying database that is able to read and write encrypted
	// values. This Database will not perform any encrypting or decrypting of
	// values and is not recommended to be used when implementing a VM.
	GetRawDatabase(username, password string) (database.Database, error)
}

type blockchainKeystore struct {
	blockchainID ids.ID
	ks           *keystore
}

func (bks *blockchainKeystore) GetDatabase(username, password string) (*encdb.Database, error) {
	bks.ks.log.Debug("Keystore: GetDatabase called with %s from %s", username, bks.blockchainID)

	return bks.ks.GetDatabase(bks.blockchainID, username, password)
}

func (bks *blockchainKeystore) GetRawDatabase(username, password string) (database.Database, error) {
	bks.ks.log.Debug("Keystore: GetRawDatabase called with %s from %s", username, bks.blockchainID)

	return bks.ks.GetRawDatabase(bks.blockchainID, username, password)
}
