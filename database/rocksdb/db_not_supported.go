// +build !linux !amd64 !rocksdballowed

// ^ Only build this file if this computer is not Linux OR it's not AMD64 OR rocksdb is not allowed
// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package rocksdb

import (
	"errors"

	"github.com/flare-foundation/flare/database"
	"github.com/flare-foundation/flare/utils/logging"
)

var errUnsupportedDatabase = errors.New("database isn't suppported")

// New returns an error.
func New(file string, log logging.Logger) (database.Database, error) {
	return nil, errUnsupportedDatabase
}
