package cli

import (
	"github.com/dplabs/cbox/internal/pkg"
	"github.com/spf13/cobra"
)

var cloudCommandsCmd = &cobra.Command{
	Use:     "commands",
	Aliases: []string{"command"},
	Args:    cobra.ExactArgs(0),
	Short:   "Cloud operations for commands",
	Long:    pkg.Logo,
}

var cloudCommandsListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l", "view"},
	Args:    cobra.ExactArgs(1),
	Short:   "List the content of a space from cbox cloud",
	Long:    pkg.Logo,
	Run:     ctrl.CloudCommandList,
}

var cloudCopyCmd = &cobra.Command{
	Use:   "copy",
	Args:  cobra.ExactArgs(2),
	Short: "Copy a remote command into a local space",
	Long:  pkg.Logo,
	Run:   ctrl.CloudCommandCopy,
}

func init() {
	cloudCmd.AddCommand(cloudCommandsCmd)
	cloudCommandsCmd.AddCommand(cloudCommandsListCmd)
	cloudCommandsCmd.AddCommand(cloudCopyCmd)
}
