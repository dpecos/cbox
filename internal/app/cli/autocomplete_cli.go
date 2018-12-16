package cli

import (
	"github.com/dplabs/cbox/internal/pkg"
	"github.com/spf13/cobra"
)

var autocompleteCmd = &cobra.Command{
	Use:   "autocomplete",
	Short: "Generates shell completion scripts",
	Args:  cobra.ExactArgs(1),
	Long:  pkg.Logo,
	Run:   ctrl.GenerateShellCompletion,
}

func init() {
	rootCmd.AddCommand(autocompleteCmd)
}
