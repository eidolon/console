#!/bin/bash

set -e

NOTES=$PWD/release-notes

version=$(cat version/version)

cat > $NOTES/notes.md <<EOF
Released v${version}!
EOF
