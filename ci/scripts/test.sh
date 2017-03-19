#!/usr/bin/env bash

set -ex

SCRIPT_PATH="$(dirname "$0")"

# Ensure we're running in the correct directory.
pushd "$SCRIPT_PATH/../../.." > /dev/null

# We'll need to set the GOPATH to the current path, then set it up.
export GOPATH=$PWD

# Move code into GOPATH
mkdir -p src/github.com/eidolon
cp -r ./console src/github/eidolon/

# Debugging Information:
whoami
env

# Pre-install
go get -u github.com/golang/lint/golint

# Install
go get -u ./...

# Script
golint -set_exit_status ./...
go vet ./...
go test -cover ./...

popd > /dev/null
