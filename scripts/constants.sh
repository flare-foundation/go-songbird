#!/usr/bin/env bash
#
# Use lower_case variables in the scripts and UPPER_CASE variables for override
# Use the constants.sh for env overrides
# Use the versions.sh to specify versions
#

AVALANCHE_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd ) # Directory above this script

# Set the PATHS
coreth_path="$AVALANCHE_PATH/coreth"

# Where AvalancheGo binary goes
build_dir="$AVALANCHE_PATH/build"
avalanchego_path="$build_dir/avalanchego"
plugin_dir="$build_dir/plugins"
evm_path="$plugin_dir/evm"

# Current branch
current_branch=$(git symbolic-ref -q --short HEAD || git describe --tags --exact-match)

git_commit=${AVALANCHEGO_COMMIT:-$( git rev-list -1 HEAD )}

# Static compilation
static_ld_flags=''
if [ "${STATIC_COMPILATION:-}" = 1 ]
then
    export CC=musl-gcc
    which $CC > /dev/null || ( echo $CC must be available for static compilation && exit 1 )
    static_ld_flags=' -extldflags "-static" -linkmode external '
fi
