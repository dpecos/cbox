#!/usr/bin/env sh

set -e
set -v

go test -json ./... > report.json
go test -coverprofile=coverage.out -v ./...
