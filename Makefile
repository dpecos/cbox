BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
COMMIT=$(shell git rev-parse HEAD)
TAG=$(shell git describe --tags 2> /dev/null)
BUILD=$(shell date +%FT%T%z)

build-dev: version
	sh ./scripts/compile.sh build $(VERSION) $(BRANCH)-$(COMMIT)@$(BUILD) dev

build-prod: version
	sh ./scripts/compile.sh build $(VERSION) $(BRANCH)-$(COMMIT)@$(BUILD) prod

install-dev: version
	sh ./scripts/compile.sh install $(VERSION) $(BRANCH)-$(COMMIT)@$(BUILD) dev

install-prod: version
	sh ./scripts/compile.sh install $(VERSION) $(BRANCH)-$(COMMIT)@$(BUILD) prod

version:
  VERSION := $(if $(TAG),$(TAG),0.0.0)

check:
	gosec ./...

test:
	sh ./scripts/test.sh