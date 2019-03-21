package cli

import (
	"github.com/dplabs/cbox/src/controllers"
	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/cobra"
)

var cloudCommandsListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Args:    cobra.ExactArgs(1),
	Short:   "Fetch a list of commands from the cloud",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.CloudCommandList(args[0]) },
}

var cloudCopyCmd = &cobra.Command{
	Use:   "copy",
	Args:  cobra.ExactArgs(2),
	Short: "Create a local copy of the commands retrieved from the cloud",
	Long:  tools.Logo,
	Run:   func(cmd *cobra.Command, args []string) { ctrl.CloudCommandCopy(args[0], args[1]) },
}

var cloudViewCmd = &cobra.Command{
	Use:     "view",
	Aliases: []string{"v"},
	Args:    cobra.ExactArgs(1),
	Short:   "Display all the details for an specific command from the cloud",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.CloudCommandView(args[0]) },
}

func init() {
	cloudCmd.AddCommand(cloudCommandsListCmd)
	cloudCmd.AddCommand(cloudCopyCmd)
	cloudCmd.AddCommand(cloudViewCmd)

	cloudCommandsListCmd.Flags().BoolVarP(&controllers.ShowCommandsSourceFlag, "view", "v", false, "Show all details about commands")

}
