#!/usr/bin/env bash
#
# Use lower_case variables in the scripts and UPPER_CASE variables for override
# Use the constants.sh for env overrides
# Use the versions.sh to specify versions
#

# Set up the versions to be used
# Don't export them as their used in the context of other calls  (old version - c0feec6aab23a5ed4bcd2551cd5b3381b7a2e16e, b4f79eab608d0ccbe06d58a52b77dfc232514028, 8933289d43ce4f129a91a7252b707d537e817274)
coreth_version=${CORETH_VERSION:-'8933289d43ce4f129a91a7252b707d537e817274'}

# Changes to the minimum golang version must also be replicated in
# README.md
# go.mod
go_version_minimum="1.16.8"
