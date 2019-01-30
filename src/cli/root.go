package cli

import (
	"log"
	"strings"

	"github.com/dplabs/cbox/src/controllers"
	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/tools"
	"github.com/dplabs/cbox/src/tools/console"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ctrl *controllers.CLIController

var rootCmd = &cobra.Command{
	Use:  "cbox",
	Long: tools.Logo,
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&console.DisableColors, "no-color", "", false, "Disable color in the output")
	rootCmd.PersistentFlags().BoolVarP(&controllers.SkipQuestionsFlag, "yes", "", false, "Answer 'yes' to any question")

	cobra.OnInitialize(func() {
		path := ""
		core.LoadSettings(path)

		cbox := core.Load()

		cloud, err := core.CloudClient()
		if err != nil {
			log.Fatal(err)
		}

		ctrl = controllers.InitController(cbox, cloud)
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if !strings.Contains(err.Error(), "unknown command") {
			log.Fatal(err)
		}
	} else {
		if ctrl != nil { // only if the config is initialized
			if err := viper.WriteConfig(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
