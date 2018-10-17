BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
COMMIT=$(shell git rev-parse HEAD)
TAG=$(shell git describe --tags 2> /dev/null)
BUILD=$(shell date +%FT%T%z)

build-dev: version
	sh ./scripts/compile.sh build $(VERSION) $(BUILD) dev

build-prod: version
	sh ./scripts/compile.sh build $(VERSION) $(BUILD) prod

install-dev: version
	sh ./scripts/compile.sh install $(VERSION) $(BUILD) dev

install-prod: version
	sh ./scripts/compile.sh install $(VERSION) $(BUILD) prod

version:
  VERSION := $(if $(TAG),$(TAG),$(BRANCH)-$(COMMIT))

check:
	gosec ./...

test:
	go test ./...
