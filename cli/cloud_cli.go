package cli

import (
	"github.com/dpecos/cbox/tools"
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

var cloudListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Args:    cobra.ExactArgs(1),
	Short:   "List the content of a space from cbox cloud",
	Long:    tools.Logo,
	Run:     ctrl.CloudCommandList,
}

func init() {
	rootCmd.AddCommand(cloudCmd)
	cloudCmd.AddCommand(cloudLoginCmd)
	cloudCmd.AddCommand(cloudLogoutCmd)
	cloudCmd.AddCommand(cloudPublishCmd)
	cloudCmd.AddCommand(cloudListCmd)

	cloudListCmd.Flags().BoolVarP(&viewSnippet, "view", "v", false, "Show all details about commands")
}
