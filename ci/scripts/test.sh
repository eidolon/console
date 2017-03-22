#!/usr/bin/env bash

set -ex

SCRIPT_PATH="$(dirname "$0")"

pushd "$SCRIPT_PATH/../../.." > /dev/null

export GOPATH=$PWD
export PATH=$PATH:$GOPATH/bin

# Move code into GOPATH
mkdir -p src/github.com/eidolon
cp -r ./console src/github.com/eidolon/console

# Debugging Information:
whoami
env | sort

# Move into source directory
pushd src/github.com/eidolon/console > /dev/null

# Pre-install
go get -u github.com/golang/lint/golint

# Install
go get ./...

# Script
golint -set_exit_status ./...
go vet ./...
go test -cover ./...

# Leave
popd > /dev/null
popd > /dev/null
