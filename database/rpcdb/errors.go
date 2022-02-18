// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcdb

import (
	"github.com/flare-foundation/flare/database"
)

var (
	errCodeToError = map[uint32]error{
		1: database.ErrClosed,
		2: database.ErrNotFound,
<<<<<<< HEAD
		3: database.ErrAvoidCorruption,
	}
	errorToErrCode = map[error]uint32{
		database.ErrClosed:          1,
		database.ErrNotFound:        2,
		database.ErrAvoidCorruption: 3,
=======
	}
	errorToErrCode = map[error]uint32{
		database.ErrClosed:   1,
		database.ErrNotFound: 2,
>>>>>>> upstream-v1.7.5
	}
)

func errorToRPCError(err error) error {
<<<<<<< HEAD
	switch err {
	case database.ErrClosed, database.ErrNotFound, database.ErrAvoidCorruption:
		return nil
	default:
		return err
	}
=======
	if _, ok := errorToErrCode[err]; ok {
		return nil
	}
	return err
>>>>>>> upstream-v1.7.5
}
