package cli

import (
	"fmt"

	"github.com/dpecos/cbox/tools/console"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) Version(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", tools.Logo)
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Version:"), core.Version)
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Build:"), core.Build)
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Cloud:"), core.CloudURL())
	fmt.Println()
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Author:"), "Daniel Pecos Martinez")
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Homepage:"), "https://cbox.dplabs.io")
}
