#!/bin/bash
# Compute "dirhash" of a go module directory.
#
# Original go code: https://cs.opensource.google/go/x/mod/+/master:sumdb/dirhash/hash.go
#
# Author: Mark Lodato <lodato@google.com>
#
# Copyright 2022 Google LLC.
# SPDX-License-Identifier: Apache-2.0

if [[ $# != 2 ]]; then
  cat >&2 <<EOF
USAGE: $(basename "$0") <directory> <modulename>@<version>
       $(basename "$0") <zipfile>

Outputs the "dirhash" of  a go module. If given a directory, the module name and
version must be provided. If given a zip file, each path in the file must start
with <modulename>@<version>.

Example:
    \$ $(basename "$0") . mymodule@v1.2.3
EOF
  exit 1
fi
DIR=$1
MODNAME=$2

if [[ ! -d "$DIR" ]]; then
  echo "ERROR: not a directory: $DIR" >&2
  exit 1
fi

if [[ -z "$MODNAME" ]]; then
  MODNAME=$DIR
elif [[ "$MODNAME" =~ [/\n] ]]; then
  echo "ERROR: module name must not contain a slash or newline: $MODNAME" >&2
  exit 1
fi

# Prints the sha256sum output of all files in $1, sorted in the C locale,
# replacing the directory portion with $2.
sha256sum_directory() {
  (
    cd "$1" && find . -type f \
      | LC_ALL=C sort \
      | xargs -r sha256sum \
      | sed "s/  ./  $2/"
  )
}

# Converts the hex on stdin, up to the first space, to base64.
hex_to_base64() {
  cut -f1 -d' ' \
    | xxd -r -p \
    | base64
}

echo "h1:$(sha256sum_directory "$DIR" "$MODNAME" | sha256sum | hex_to_base64)"