package cli

import (
	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Args:  cobra.ExactArgs(0),
	Short: "Display current cbox version",
	Run:   ctrl.Version,
	Long:  tools.Logo,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
