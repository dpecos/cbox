package cli

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Args:  cobra.MaximumNArgs(0),
	Short: "Display current cbox version",
	Run:   ctrl.Version,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
