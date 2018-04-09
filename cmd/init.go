package cmd

import (
	"github.com/dpecos/cmdbox/db"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your cmdbox",
	Long: `It creates a storage for your cmbdbox.

WARNING! It will remove any previous existing database that could exist in the same path.`,
	Run: func(cmd *cobra.Command, args []string) {
		db.Init(dbPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
