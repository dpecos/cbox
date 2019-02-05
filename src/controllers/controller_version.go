package controllers

import (
	"github.com/dplabs/cbox/src/tools"
	"github.com/dplabs/cbox/src/tools/tty"
)

func (ctrl *CLIController) Version() {
	tty.Print("%s\n", tools.Logo)
	tty.Print("%s %s\n", tty.ColorBoldBlack("Version:"), ctrl.cbox.Version)
	tty.Print("%s %s\n", tty.ColorBoldBlack("Build:"), ctrl.cbox.Build)
	tty.Print("%s %s\n", tty.ColorBoldBlack("Cloud Environment:"), ctrl.cloud.Environment)
	tty.Print("\n")
	tty.Print("%s %s\n", tty.ColorBoldBlack("Author:"), "Daniel Pecos Martinez")
	tty.Print("%s %s\n", tty.ColorBoldBlack("Homepage:"), "https://cbox.dplabs.io")
}
