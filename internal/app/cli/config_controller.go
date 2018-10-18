package cli

import (
	"github.com/dpecos/cbox/internal/pkg"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

func (ctrl *CLIController) ConfigSet(cmd *cobra.Command, args []string) {
	config := args[0]
	value := args[1]
	viper.Set(config, value)
	pkg.PrintSetting(config, value)
}

func (ctrl *CLIController) ConfigGet(cmd *cobra.Command, args []string) {
	config := args[0]
	value := viper.GetString(config)
	pkg.PrintSetting(config, value)
}
