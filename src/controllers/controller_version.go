package controllers

import (
	"fmt"

	"github.com/dplabs/cbox/src/tools"
	"github.com/dplabs/cbox/src/tools/tty"
)

func (ctrl *CLIController) Version() {
	fmt.Printf("%s\n", tools.Logo)
	fmt.Printf("%s %s\n", tty.ColorBoldBlack("Version:"), ctrl.cbox.Version)
	fmt.Printf("%s %s\n", tty.ColorBoldBlack("Build:"), ctrl.cbox.Build)
	fmt.Printf("%s %s\n", tty.ColorBoldBlack("Cloud Environment:"), ctrl.cloud.Environment)
	fmt.Println()
	fmt.Printf("%s %s\n", tty.ColorBoldBlack("Author:"), "Daniel Pecos Martinez")
	fmt.Printf("%s %s\n", tty.ColorBoldBlack("Homepage:"), "https://cbox.dplabs.io")
}
