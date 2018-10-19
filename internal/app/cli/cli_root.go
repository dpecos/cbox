package cli

import (
	"log"
	"os"

	"github.com/dpecos/cbox/internal/app/core"
	"github.com/dpecos/cbox/internal/pkg"
	"github.com/dpecos/cbox/pkg/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cboxInstance *models.CBox

var rootCmd = &cobra.Command{
	Use:   "cbox",
	Short: "",
	Long:  pkg.Logo,
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash completion scripts",
	Long: `To load completion run

. <(bitbucket completion)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(bitbucket completion)
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := rootCmd.GenBashCompletion(os.Stdout); err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	if err := viper.WriteConfig(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(func() {
		path := ""
		core.LoadSettings(path)
		cboxInstance = core.Load()
	})

	rootCmd.AddCommand(completionCmd)
}
