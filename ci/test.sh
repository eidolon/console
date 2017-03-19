#!/usr/bin/env bash

SCRIPT_PATH="$(dirname "$0")"

pushd "$SCRIPT_PATH/.." > /dev/null

set -ex

# Pre-install
go get -u github.com/golang/lint/golint

# Install
go get -u ./...

# Script
golint -set_exit_status ./...
go vet ./...
go test -cover ./...

popd > /dev/null
