#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Avalanchego root folder
FLARE_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
CORETH_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd ../coreth && pwd )
# Load the versions
source "$FLARE_PATH"/scripts/versions.sh
# Load the constants
source "$FLARE_PATH"/scripts/constants.sh

mkdir -p $GOPATH/pkg/mod/github.com/flare-foundation/flare@$flare_version
rsync -ar --delete $FLARE_PATH/* $GOPATH/pkg/mod/github.com/flare-foundation/flare@$flare_version
 
# Download dependencies
echo "Downloading dependencies..."
go mod download -modcacherw

# Build flare
"$FLARE_PATH"/scripts/build_flare.sh

# Build coreth
"$FLARE_PATH"/scripts/build_coreth.sh

# Exit build successfully if the binaries are created
if [[ -f "$flare_path" && -f "$evm_path" ]]; then
        echo "Build Successful"
        exit 0
else
        echo "Build failure" >&2
        exit 1
fi
