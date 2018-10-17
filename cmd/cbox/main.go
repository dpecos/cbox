package main

import (
	"log"

	"github.com/dpecos/cbox/internal/cli"
	"github.com/dpecos/cbox/tools/console"
)

func main() {
	// log.SetPrefix("\x1b[31;1mcbox error: ")
	log.SetPrefix(console.ColorBoldRed("\nError:") + " ")
	log.SetFlags(0)

	cli.Execute()
}
