package cli

import (
	"github.com/dplabs/cbox/src/controllers"
	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/cobra"
)

var cloudCommandsCmd = &cobra.Command{
	Use:     "commands",
	Aliases: []string{"command"},
	Args:    cobra.ExactArgs(0),
	Short:   "Cloud operations for commands",
	Long:    tools.Logo,
}

var cloudCommandsListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l", "view"},
	Args:    cobra.ExactArgs(1),
	Short:   "List the content of a space from cbox cloud",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.CloudCommandList(args) },
}

var cloudCopyCmd = &cobra.Command{
	Use:   "copy",
	Args:  cobra.ExactArgs(2),
	Short: "Copy a remote command into a local space",
	Long:  tools.Logo,
	Run:   func(cmd *cobra.Command, args []string) { ctrl.CloudCommandCopy(args) },
}

func init() {
	cloudCmd.AddCommand(cloudCommandsCmd)
	cloudCommandsCmd.AddCommand(cloudCommandsListCmd)
	cloudCommandsCmd.AddCommand(cloudCopyCmd)

	cloudCommandsListCmd.Flags().BoolVarP(&controllers.ShowCommandsSourceFlag, "view", "v", false, "Show all details about commands")

}
