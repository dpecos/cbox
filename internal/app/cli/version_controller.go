package cli

import (
	"fmt"

	"github.com/dpecos/cbox/internal/app/core"
	"github.com/dpecos/cbox/internal/pkg"
	"github.com/dpecos/cbox/internal/pkg/console"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) Version(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", pkg.Logo)
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Version:"), core.Version)
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Build:"), core.Build)
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Cloud:"), core.CloudURL())
	fmt.Println()
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Author:"), "Daniel Pecos Martinez")
	fmt.Printf("%s %s\n", console.ColorBoldBlack("Homepage:"), "https://cbox.dplabs.io")
}
