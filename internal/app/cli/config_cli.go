package cli

import (
	"github.com/dplabs/cbox/internal/pkg"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Args:  cobra.ExactArgs(0),
	Short: "Manage cbox settings",
	Long:  pkg.Logo,
}

var setConfigCmd = &cobra.Command{
	Use:   "set",
	Args:  cobra.ExactArgs(2),
	Short: "Set the value for a config setting",
	Run:   ctrl.ConfigSet,
}

var getConfigCmd = &cobra.Command{
	Use:   "get",
	Args:  cobra.ExactArgs(1),
	Short: "Get the value for a config setting",
	Run:   ctrl.ConfigGet,
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(setConfigCmd)
	configCmd.AddCommand(getConfigCmd)
}
