package cli

import (
	"github.com/dpecos/cbox/internal/pkg"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Args:  cobra.ExactArgs(0),
	Short: "Display current cbox version",
	Run:   ctrl.Version,
	Long:  pkg.Logo,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
