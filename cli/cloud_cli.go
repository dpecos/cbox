package cli

import (
	"github.com/spf13/cobra"
)

var cloudCmd = &cobra.Command{
	Use:   "cloud",
	Args:  cobra.ExactArgs(0),
	Short: "Login & sync your spaces to the cloud",
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

var cloudPublishCmd = &cobra.Command{
	Use:   "publish",
	Args:  cobra.ExactArgs(1),
	Short: "Publish an space",
	Run:   ctrl.CloudPublishSpace,
}

func init() {
	rootCmd.AddCommand(cloudCmd)
	cloudCmd.AddCommand(cloudLoginCmd)
	cloudCmd.AddCommand(cloudLogoutCmd)
	cloudCmd.AddCommand(cloudPublishCmd)
}
