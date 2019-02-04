package controllers

import (
	"fmt"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/tools"
	"github.com/dplabs/cbox/src/tools/console"
)

func (ctrl *CLIController) Version() {
	fmt.Printf("%s\n", tools.Logo)
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Version:"), core.Version)
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Build:"), core.Build)
	fmt.Println()
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Author:"), "Daniel Pecos Martinez")
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Homepage:"), "https://cbox.dplabs.io")
}
