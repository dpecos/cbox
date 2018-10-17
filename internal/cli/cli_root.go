package cli

import (
	"log"
	"os"

	"github.com/dpecos/cbox/internal/core"
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "cbox",
	Short: "",
	Long:  tools.Logo,
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
		rootCmd.GenBashCompletion(os.Stdout)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	viper.WriteConfig()
}

func init() {
	cobra.OnInitialize(func() {
		core.InitCBox("")
	})

	rootCmd.AddCommand(completionCmd)
}
