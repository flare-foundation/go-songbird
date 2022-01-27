#!/usr/bin/env bash
#
# Use lower_case variables in the scripts and UPPER_CASE variables for override
# Use the constants.sh for env overrides
# Use the versions.sh to specify versions
#

# Set up the versions to be used
# Don't export them as their used in the context of other calls
coreth_version=${CORETH_VERSION:-'3cdb42a322a009a36a165122b21e2b72fab82635'}

# Changes to the minimum golang version must also be replicated in
# README.md
# go.mod
go_version_minimum="1.16.8"
