package cli

import (
	"fmt"

	"github.com/dpecos/cbox/tools"
	"github.com/spf13/cobra"
)

var cboxVersion = "development"

func (ctrl *CLIController) Version(cmd *cobra.Command, args []string) {
	fmt.Printf("%s", tools.Logo)
	fmt.Printf("Version: %s\n", cboxVersion)
	fmt.Printf("Homepage: https://cbox.dplabs.io\n")
	fmt.Printf("Author: Daniel Pecos Martinez\n")
}
