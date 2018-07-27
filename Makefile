BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
COMMIT=$(shell git rev-parse HEAD)
TAG=$(shell git describe --tags 2> /dev/null)
BUILD=$(shell date +%FT%T%z)

build-dev: version
	go build -ldflags "-X github.com/dpecos/cbox/core.Version=$(VERSION) -X github.com/dpecos/cbox/core.Build=$(BUILD)"

build-prod: version
	go build -ldflags "-X github.com/dpecos/cbox/core.Version=$(VERSION) -X github.com/dpecos/cbox/core.Build=$(BUILD) -X github.com/dpecos/cbox/core.Env=prod"

install-dev: version
	go install -ldflags "-X github.com/dpecos/cbox/core.Version=$(VERSION) -X github.com/dpecos/cbox/core.Build=$(BUILD)"

install-prod: version
	go install -ldflags "-X github.com/dpecos/cbox/core.Version=$(VERSION) -X github.com/dpecos/cbox/core.Build=$(BUILD) -X github.com/dpecos/cbox/core.Env=prod"

version:
  VERSION := $(if $(TAG),$(TAG),$(BRANCH)-$(COMMIT))

test:
	go test ./...
