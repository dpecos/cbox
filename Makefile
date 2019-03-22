BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
COMMIT=$(shell git rev-parse HEAD)
TAG=$(shell git describe --tags 2> /dev/null)
BUILD=$(shell date +%FT%T%z)

build: version
	sh ./scripts/compile.sh build $(VERSION) $(BRANCH)-$(COMMIT)@$(BUILD) prod

build-test: version
	sh ./scripts/compile.sh build $(VERSION) $(BRANCH)-$(COMMIT)@$(BUILD) test

install: version
	sh ./scripts/compile.sh install $(VERSION) $(BRANCH)-$(COMMIT)@$(BUILD) prod

install-test: version
	sh ./scripts/compile.sh install $(VERSION) $(BRANCH)-$(COMMIT)@$(BUILD) test

version:
  VERSION := $(if $(TAG),$(TAG),0.0.0)

check:
	gosec ./...

test:
	sh ./scripts/test.sh

release:
	. ./scripts/pre-release.sh && goreleaser --rm-dist	

snapshot-release:
	. ./scripts/pre-release.sh && goreleaser --rm-dist --snapshot