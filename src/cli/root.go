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
	rootCmd.PersistentFlags().BoolVar(&tty.DisableColors, "no-color", false, "Disable color in the output")
	rootCmd.PersistentFlags().BoolVar(&tty.DisableOutput, "silent", false, "Completely disable any output")
	rootCmd.PersistentFlags().BoolVar(&controllers.SkipQuestionsFlag, "yes", false, "Answer 'yes' to any question")
	rootCmd.PersistentFlags().StringVarP(&controllers.ListingsModeOption, "listings-mode", "m", "", "Use 'fzf' (interactive) to interact with commands listings or just print them as an static list (static)")
	rootCmd.PersistentFlags().StringVarP(&controllers.ListingsSortOption, "listings-sort", "s", "", "Sort commands listings by name (default) or date")
}

func loadParameterValuesFromConfig() {
	if controllers.ListingsModeOption == "" {
		controllers.ListingsModeOption = viper.GetString("cbox.results.mode")
	}
	if controllers.ListingsSortOption == "" {
		controllers.ListingsSortOption = viper.GetString("cbox.results.sort")
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
