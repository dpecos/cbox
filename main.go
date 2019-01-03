package main

import (
	"log"

	"github.com/dplabs/cbox/src/cli"
	"github.com/dplabs/cbox/src/tools/console"
)

func main() {
	log.SetPrefix(console.ColorBoldRed("\nError:") + " ")
	log.SetFlags(0)

	cli.Execute()
}
