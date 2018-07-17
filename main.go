package main

import (
	"log"

	"github.com/dpecos/cbox/cli"
	"github.com/dpecos/cbox/tools/console"
)

func main() {
	// log.SetPrefix("\x1b[31;1mcbox error: ")
	log.SetPrefix(console.ColorBgBoldRed("Error:") + " ")
	log.SetFlags(0)

	cli.Execute()
}
