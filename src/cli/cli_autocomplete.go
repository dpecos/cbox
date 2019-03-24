package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/cobra"
)

var autocompleteCmd = &cobra.Command{
	Use:   "autocomplete",
	Short: "Generate an auto-completion script for your shell",
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
			log.Fatal(fmt.Sprintf("Shell '%s' not supported", shell))
		}
	},
}

func init() {
	rootCmd.AddCommand(autocompleteCmd)
}
