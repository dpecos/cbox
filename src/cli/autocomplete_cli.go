package cli

import (
	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/cobra"
)

var autocompleteCmd = &cobra.Command{
	Use:   "autocomplete",
	Short: "Generates shell completion scripts",
	Args:  cobra.ExactArgs(1),
	Long:  tools.Logo,
	Run:   ctrl.GenerateShellCompletion,
}

func init() {
	rootCmd.AddCommand(autocompleteCmd)
}
