build:
	go build -ldflags "-X github.com/dpecos/cbox/cli.cboxVersion=$(shell git rev-parse HEAD)"