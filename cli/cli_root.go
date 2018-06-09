package cli

import (
	"log"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/cobra"
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
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	core.CheckCboxDir()
}
