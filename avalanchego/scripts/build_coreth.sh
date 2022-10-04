#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Directory above this script
FLARE_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )

# Load the versions
source "$FLARE_PATH"/scripts/versions.sh

# Load the constants
source "$FLARE_PATH"/scripts/constants.sh

# Build Coreth
echo "Building Coreth @ ${coreth_version}..."
cd "$coreth_path"
go build -modcacherw -ldflags "-X github.com/flare-foundation/flare/coreth/plugin/evm.Version=$coreth_version $static_ld_flags" -o "$evm_path" "plugin/"*.go
cd "$FLARE_PATH"

# Building coreth + using go get can mess with the go.mod file.
go mod tidy
