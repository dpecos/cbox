package cli

import (
	"github.com/dplabs/cbox/src/controllers"
	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Args:  cobra.MinimumNArgs(1),
	Short: "Search commands matching a criteria in a given local space",
	Long:  tools.Logo,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			ctrl.SearchCommands(nil, args[0])
		} else {
			ctrl.SearchCommands(&args[0], args[1])
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().BoolVarP(&controllers.ShowCommandsSourceFlag, "view", "v", false, "Show all details about commands")
}
