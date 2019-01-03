package cli

import (
	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/cobra"
)

var cloudCmd = &cobra.Command{
	Use:   "cloud",
	Args:  cobra.ExactArgs(0),
	Short: "Login & sync your spaces to the cloud",
	Long:  tools.Logo,
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

func init() {
	rootCmd.AddCommand(cloudCmd)
	cloudCmd.AddCommand(cloudLoginCmd)
	cloudCmd.AddCommand(cloudLogoutCmd)

	cloudCommandsCmd.Flags().BoolVarP(&viewSnippet, "view", "v", false, "Show all details about commands")
}
