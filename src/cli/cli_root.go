package cli

import (
	"log"
	"strings"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools"
	"github.com/dplabs/cbox/src/tools/console"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cboxInstance *models.CBox
)

var rootCmd = &cobra.Command{
	Use:  "cbox",
	Long: tools.Logo,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if !strings.Contains(err.Error(), "unknown command") {
			log.Fatal(err)
		}
	} else {
		if cboxInstance != nil { // only if the config is initialized
			if err := viper.WriteConfig(); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&console.DisableColors, "no-color", "", false, "Disable color in the output")

	cobra.OnInitialize(func() {
		path := ""
		core.LoadSettings(path)
		cboxInstance = core.Load()
	})
}
