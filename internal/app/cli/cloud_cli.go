package cli

import (
	"github.com/dplabs/cbox/internal/pkg"
	"github.com/spf13/cobra"
)

var cloudCmd = &cobra.Command{
	Use:   "cloud",
	Args:  cobra.ExactArgs(0),
	Short: "Login & sync your spaces to the cloud",
	Long:  pkg.Logo,
}

var cloudLoginCmd = &cobra.Command{
	Use:   "login",
	Args:  cobra.ExactArgs(0),
	Short: "Login to cbox cloud using your Github account",
	Run:   ctrl.CloudLogin,
}

var cloudLogoutCmd = &cobra.Command{
	Use:   "logout",
	Args:  cobra.ExactArgs(0),
	Short: "Logout from cbox cloud",
	Run:   ctrl.CloudLogout,
}

var cloudCommandsCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"c", "cmd", "commands", "ls", "l"},
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
	rootCmd.AddCommand(cloudCmd)
	cloudCmd.AddCommand(cloudLoginCmd)
	cloudCmd.AddCommand(cloudLogoutCmd)
	cloudCmd.AddCommand(cloudCommandsCmd)
	cloudCmd.AddCommand(cloudCopyCmd)

	cloudCommandsCmd.Flags().BoolVarP(&viewSnippet, "view", "v", false, "Show all details about commands")
}
