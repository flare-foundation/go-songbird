// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package version

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func customApricotPhaseTime(phase uint) (time.Time, bool) {
	key := fmt.Sprintf("CUSTOM_APRICOT_PHASE_%d_TIME", phase)
	value := os.Getenv(key)
	if value == "" {
		return time.Time{}, false
	}
	seconds, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("unix timestamp for %s is not an integer: %s", key, err))
	}
	timestamp := time.Unix(seconds, 0)

	return timestamp, true
}
