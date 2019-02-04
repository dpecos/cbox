package main

import (
	"log"

	"github.com/dplabs/cbox/src/cli"
	"github.com/dplabs/cbox/src/tools/tty"
)

func main() {
	log.SetPrefix(tty.ColorBoldRed("\nError:") + " ")
	log.SetFlags(0)

	cli.Execute()
}
