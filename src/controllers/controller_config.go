package controllers

import (
	"github.com/dplabs/cbox/src/tools/console"
	"github.com/spf13/viper"
)

func (ctrl *CLIController) ConfigSet(config string, value string) {
	viper.Set(config, value)
	console.PrintSetting(config, value)
}

func (ctrl *CLIController) ConfigGet(config string) {
	value := viper.GetString(config)
	console.PrintSetting(config, value)
}
