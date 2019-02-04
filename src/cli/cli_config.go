package cli

import (
	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Args:  cobra.ExactArgs(0),
	Short: "Manage cbox settings",
	Long:  tools.Logo,
}

var setConfigCmd = &cobra.Command{
	Use:   "set",
	Args:  cobra.ExactArgs(2),
	Short: "Set the value for a config setting",
	Run:   func(cmd *cobra.Command, args []string) { ctrl.ConfigSet(args[0], args[1]) },
}

var getConfigCmd = &cobra.Command{
	Use:   "get",
	Args:  cobra.ExactArgs(1),
	Short: "Get the value for a config setting",
	Run:   func(cmd *cobra.Command, args []string) { ctrl.ConfigGet(args[0]) },
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(setConfigCmd)
	configCmd.AddCommand(getConfigCmd)
}
