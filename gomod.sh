#!/bin/bash
set -eu

if [ -f ./go.mod ]; then
    exit 0
fi

touch go.mod

PROJECT_NAME=$(basename $(pwd | xargs dirname))
CURRENT_DIR=$(basename $(pwd))

CONTENT=$(cat <<-EOD
module github.com/${PROJECT_NAME}/${CURRENT_DIR}
EOD
)

echo "$CONTENT" > go.mod
