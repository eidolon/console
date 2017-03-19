#!/usr/bin/env bash

SCRIPT_PATH="$(dirname "$0")"

pushd "$SCRIPT_PATH" > /dev/null

fly -t ci set-pipeline -p eidolon:console -c pipeline.yml -l variables.yml

popd > /dev/null
