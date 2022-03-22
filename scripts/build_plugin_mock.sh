#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Directory above this script
FLARE_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )

# Load the constants
source "$FLARE_PATH"/scripts/constants.sh

# Build VM Plugin Mock
echo "Building VM plugin mock..."
cd "$FLARE_PATH"/testing/mock
go build -ldflags "$static_ld_flags" -o "$plugin_dir"/mock "plugin/"*.go
cd "$FLARE_PATH"

# Building coreth + using go get can mess with the go.mod file.
go mod tidy
