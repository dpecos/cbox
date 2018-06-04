package cli

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Args:  cobra.ExactArgs(0),
	Short: "Initialize your cmdbox",
	Long: `It creates a storage for your cmbdbox.

WARNING! It will remove any previous existing database that could exist in the same path.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
