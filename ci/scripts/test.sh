#!/usr/bin/env bash

set -ex

export GOPATH="$PWD/go"
export PATH="$PATH:$GOPATH/bin"

pushd "$(dirname "$0")/../.." > /dev/null
    # Prepare
    go get -u github.com/golang/lint/golint
    go get -d -t ./...

    # Run
    golint -set_exit_status ./...
    go vet ./...
    go test -cover ./...
popd > /dev/null
