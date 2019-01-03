package cli

import (
	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

func (ctrl *CLIController) ConfigSet(cmd *cobra.Command, args []string) {
	config := args[0]
	value := args[1]
	viper.Set(config, value)
	tools.PrintSetting(config, value)
}

func (ctrl *CLIController) ConfigGet(cmd *cobra.Command, args []string) {
	config := args[0]
	value := viper.GetString(config)
	tools.PrintSetting(config, value)
}
