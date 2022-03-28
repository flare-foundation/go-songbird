// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposer

import (
	"bytes"

	"github.com/flare-foundation/flare/snow/validation"
)

type sortByID []validation.Validator

func (s sortByID) Len() int {
	return len(s)
}
func (s sortByID) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortByID) Less(i, j int) bool {
	return bytes.Compare(s[i].ID().Bytes(), s[j].ID().Bytes()) < 0
}
