#!/usr/bin/env bash

set -ex

SCRIPT_PATH="$(dirname "$0")"

pushd "$SCRIPT_PATH/../.." > /dev/null

# Pre-install
go get -u github.com/golang/lint/golint

# Install
go get -d -t ./...

# Script
golint -set_exit_status ./...
go vet ./...
go test -cover ./...

# Leave
popd > /dev/null
