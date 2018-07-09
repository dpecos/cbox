package main

import (
	"log"

	"github.com/dpecos/cbox/cli"
	"github.com/logrusorgru/aurora"
)

func main() {
	// log.SetPrefix("\x1b[31;1mcbox error: ")
	log.SetPrefix(aurora.BgRed(aurora.Bold("Error:")).String() + " ")
	log.SetFlags(0)

	cli.Execute()
}
