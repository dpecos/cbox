#!/usr/bin/env sh

set -e
set -x

COMMAND=$1
VERSION=$2
BUILD=$3
ENV=$4

go $COMMAND \
    -ldflags "-X github.com/dplabs/cbox/src/core.Version=$VERSION -X github.com/dplabs/cbox/src/core.Build=$BUILD -X github.com/dplabs/cbox/src/core.Env=$ENV"
