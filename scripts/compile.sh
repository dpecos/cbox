#!/usr/bin/env sh

COMMAND=$1
VERSION=$2
BUILD=$3
ENV=$4

go $COMMAND \
    -ldflags "-X github.com/dpecos/cbox/internal/app/core.Version=$VERSION -X github.com/dpecos/cbox/internal/app/core.Build=$BUILD -X github.com/dpecos/cbox/internal/app/core.Env=$ENV" \
    ./...