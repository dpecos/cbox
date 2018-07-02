COMMIT=$(shell git rev-parse HEAD)
BUILD=$(shell date +%FT%T%z)

build:
	go build -ldflags "-X github.com/dpecos/cbox/cli.cboxVersion=$(COMMIT) -X github.com/dpecos/cbox/cli.cboxBuild=$(BUILD)"