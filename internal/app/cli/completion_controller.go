package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/dpecos/cbox/internal/pkg/console"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) GenerateShellCompletion(cmd *cobra.Command, args []string) {
	shell := args[0]

	switch shell {
	case "bash":
		if err := rootCmd.GenBashCompletion(os.Stdout); err != nil {
			log.Fatal(err)
		}
	case "zsh":
		if err := rootCmd.GenZshCompletion(os.Stdout); err != nil {
			log.Fatal(err)
		}
	default:
		console.PrintError(fmt.Sprintf("Shell '%s' not supported", shell))
	}
}
