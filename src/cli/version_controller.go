package cli

import (
	"fmt"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/tools"
	"github.com/dplabs/cbox/src/tools/console"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) Version(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", tools.Logo)
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Version:"), core.Version)
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Build:"), core.Build)
	fmt.Println()
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Author:"), "Daniel Pecos Martinez")
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Homepage:"), "https://cbox.dplabs.io")
}
