package main

import (
	"log"

	"github.com/dpecos/cbox/internal/app/cli"
	"github.com/dpecos/cbox/internal/pkg/console"
)

func main() {
	log.SetPrefix(console.ColorBoldRed("\nError:") + " ")
	log.SetFlags(0)

	cli.Execute()
}
