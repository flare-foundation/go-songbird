#!/usr/bin/env bash
#
# Use lower_case variables in the scripts and UPPER_CASE variables for override
# Use the constants.sh for env overrides
# Use the versions.sh to specify versions
#

# Set up the versions to be used
# Don't export them as their used in the context of other calls
coreth_version=${CORETH_VERSION:-'32b0e385961b6f7eaa1352db0626bed32c7db923'}

# Changes to the minimum golang version must also be replicated in
# README.md
# go.mod
go_version_minimum="1.16.8"
