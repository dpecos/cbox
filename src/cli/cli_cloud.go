package cli

import (
	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/cobra"
)

var cloudCmd = &cobra.Command{
	Use:   "cloud",
	Args:  cobra.ExactArgs(0),
	Short: "Discover and share usefull commands with cbox's community",
	Long:  tools.Logo,
}

var cloudLoginCmd = &cobra.Command{
	Use:   "login",
	Args:  cobra.ExactArgs(0),
	Short: "Login to the cloud using your Github account",
	Run:   func(cmd *cobra.Command, args []string) { ctrl.CloudLogin() },
}

var cloudLogoutCmd = &cobra.Command{
	Use:   "logout",
	Args:  cobra.ExactArgs(0),
	Short: "Logout from the cloud - we'll miss you :'(",
	Run:   func(cmd *cobra.Command, args []string) { ctrl.CloudLogout() },
}

func init() {
	rootCmd.AddCommand(cloudCmd)
	cloudCmd.AddCommand(cloudLoginCmd)
	cloudCmd.AddCommand(cloudLogoutCmd)
}
