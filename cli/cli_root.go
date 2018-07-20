package cli

import (
	"log"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "cbox",
	Short: "",
	Long:  tools.Logo,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	viper.WriteConfig()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	cboxPath := core.CheckCboxDir("")

	viper.AddConfigPath(cboxPath)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
