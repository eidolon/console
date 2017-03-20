#!/usr/bin/env bash

set -e

SCRIPT_PATH="$(dirname "$0")"

if test ! $(which yft); then
    echo "Installing yft..."
    go get -u github.com/SeerUK/yft/...
fi

pushd "$SCRIPT_PATH" > /dev/null

# Create pipeline file, with variables substituted.
yft pipeline.tpl.yml < variables.yml > pipeline.yml

# Update pipeline in Concourse
fly -t ci set-pipeline -p eidolon:console -c pipeline.yml

popd > /dev/null
