package main

import (
	"log"

	"github.com/dpecos/cbox/cli"
)

func main() {
	log.SetPrefix("cbox error: ")
	log.SetFlags(0)

	cli.Execute()
}
