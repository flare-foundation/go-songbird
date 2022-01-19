#!/usr/bin/env bash
#
# Use lower_case variables in the scripts and UPPER_CASE variables for override
# Use the constants.sh for env overrides
# Use the versions.sh to specify versions
#

# Set up the versions to be used
# Don't export them as their used in the context of other calls  (old version - c0feec6aab23a5ed4bcd2551cd5b3381b7a2e16e, b4f79eab608d0ccbe06d58a52b77dfc232514028, 8933289d43ce4f129a91a7252b707d537e817274, 95bcf9123cbdb1f851753ce50772346e5b64e239, 9c5376022d506ff3a9bcb289755bb9c3faeb7cfb, e03971d0245505dea0cafb12c634bc5bd2819259, 4cffe7d5984c01035b5d2eb00c0a7b7b5bcd2e23, 62d7521a3cb3e4658480471dd93977d356679d16, 1e4fce4703fe12254aa2e54fd58c83df433bc1c7, 81e75ef1a37f4ac0e760973e99a7365a324f8d05, c85c66f1bbd766df4b1d7f29038b16767d080d50, e70208fd2e6fb4c3a380244905033f93acc5604a)
coreth_version=${CORETH_VERSION:-'e70208fd2e6fb4c3a380244905033f93acc5604a'}

# Changes to the minimum golang version must also be replicated in
# README.md
# go.mod
go_version_minimum="1.16.8"
