package cli

import (
	"log"
	"strings"

	"github.com/dplabs/cbox/src/controllers"
	"github.com/dplabs/cbox/src/tools"
	"github.com/dplabs/cbox/src/tools/tty"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ctrl *controllers.CLIController

var rootCmd = &cobra.Command{
	Use:  "cbox",
	Long: tools.Logo,
}

func init() {
	setupParameters()
	cobra.OnInitialize(func() {
		ctrl = controllers.InitController("")
		loadParameterValuesFromConfig()
	})
}

func setupParameters() {
	rootCmd.PersistentFlags().BoolVarP(&tty.DisableColors, "no-color", "", false, "Disable color in the output")
	rootCmd.PersistentFlags().BoolVarP(&tty.DisableOutput, "silent", "", false, "Completely disable any output")
	rootCmd.PersistentFlags().BoolVarP(&controllers.SkipQuestionsFlag, "yes", "", false, "Answer 'yes' to any question")
	rootCmd.PersistentFlags().BoolVarP(&controllers.InteractiveListingFlag, "interactive", "i", false, "Use 'fzf' to interact with command listings")
}

func loadParameterValuesFromConfig() {
	if !controllers.InteractiveListingFlag {
		controllers.InteractiveListingFlag = viper.GetBool("cbox.interactive")
	}
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
