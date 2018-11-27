package main

import (
	"log"

	"github.com/dplabs/cbox/internal/app/cli"
	"github.com/dplabs/cbox/internal/pkg/console"
)

func main() {
	log.SetPrefix(console.ColorBoldRed("\nError:") + " ")
	log.SetFlags(0)

	cli.Execute()
}
