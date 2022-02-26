#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Avalanchego root folder
FLARE_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
# Load the versions
source "$FLARE_PATH"/scripts/versions.sh
# Load the constants
source "$FLARE_PATH"/scripts/constants.sh

# Download dependencies
echo "Downloading dependencies..."
go mod download

#go get github.com/flare-foundation/flare/database/rocksdb

#build rocksdb from source, replacing broken grocksdb build script

GROCKS_DIR=$(ls $HOME/go/pkg/mod/github.com/xrpdevs/)

cd "$HOME/go/pkg/mod/github.com/xrpdevs/$GROCKS_DIR/"

if [ -z ${ROCKSDBALLOWED+x} ]; then
  echo "Building rocksdb from source, this could take some time. You must be able to use sudo to install rocksdb."
  sh "$FLARE_PATH/scripts/build_rocksdb.sh"
fi

cd "$FLARE_PATH"

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
