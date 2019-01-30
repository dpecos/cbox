package controllers

import (
	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/viper"
)

func (ctrl *CLIController) ConfigSet(args []string) {
	config := args[0]
	value := args[1]
	viper.Set(config, value)
	tools.PrintSetting(config, value)
}

func (ctrl *CLIController) ConfigGet(args []string) {
	config := args[0]
	value := viper.GetString(config)
	tools.PrintSetting(config, value)
}
