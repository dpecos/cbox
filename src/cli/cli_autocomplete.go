package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/dplabs/cbox/src/tools"
	"github.com/dplabs/cbox/src/tools/console"
	"github.com/spf13/cobra"
)

var autocompleteCmd = &cobra.Command{
	Use:   "autocomplete",
	Short: "Generates shell completion scripts",
	Args:  cobra.ExactArgs(1),
	Long:  tools.Logo,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {
	rootCmd.AddCommand(autocompleteCmd)
}
