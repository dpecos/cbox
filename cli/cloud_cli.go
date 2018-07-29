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

var cloudSpacePublishCmd = &cobra.Command{
	Use:     "publish",
	Aliases: []string{"push"},
	Args:    cobra.ExactArgs(1),
	Short:   "Publish an space",
	Run:     ctrl.CloudSpacePublish,
}

var cloudSpacePullCmd = &cobra.Command{
	Use:   "pull",
	Args:  cobra.ExactArgs(1),
	Short: "Pull latest changes of a cloud space",
	Run:   ctrl.CloudSpacePull,
}

var cloudSpaceCloneCmd = &cobra.Command{
	Use:   "clone",
	Args:  cobra.ExactArgs(1),
	Short: "Clone an space locally",
	Run:   ctrl.CloudSpaceClone,
}

var cloudCommandsCmd = &cobra.Command{
	Use:     "commands",
	Aliases: []string{"c", "cmd", "list", "ls", "l"},
	Args:    cobra.ExactArgs(1),
	Short:   "List the content of a space from cbox cloud",
	Long:    tools.Logo,
	Run:     ctrl.CloudCommandList,
}

var cloudCopyCmd = &cobra.Command{
	Use:   "copy",
	Args:  cobra.ExactArgs(2),
	Short: "Copy a remote command into a local space",
	Long:  tools.Logo,
	Run:   ctrl.CloudCommandCopy,
}

func init() {
	rootCmd.AddCommand(cloudCmd)
	cloudCmd.AddCommand(cloudLoginCmd)
	cloudCmd.AddCommand(cloudLogoutCmd)
	cloudCmd.AddCommand(cloudSpacePublishCmd)
	cloudCmd.AddCommand(cloudSpacePullCmd)
	cloudCmd.AddCommand(cloudSpaceCloneCmd)
	cloudCmd.AddCommand(cloudCommandsCmd)
	cloudCmd.AddCommand(cloudCopyCmd)

	cloudCommandsCmd.Flags().BoolVarP(&viewSnippet, "view", "v", false, "Show all details about commands")
}
