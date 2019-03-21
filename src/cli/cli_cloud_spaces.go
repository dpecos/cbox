package cli

import (
	"github.com/dplabs/cbox/src/controllers"
	"github.com/spf13/cobra"
)

var cloudSpaceInfoCmd = &cobra.Command{
	Use:   "info",
	Args:  cobra.ExactArgs(1),
	Short: "Retrieve all the details of a cloud space",
	Run:   func(cmd *cobra.Command, args []string) { ctrl.CloudSpaceInfo(args[0]) },
}

var cloudSpacePublishCmd = &cobra.Command{
	Use:     "publish",
	Aliases: []string{"push"},
	Args:    cobra.ExactArgs(1),
	Short:   "Share a local space and all its content",
	Run:     func(cmd *cobra.Command, args []string) { ctrl.CloudSpacePublish(args[0]) },
}

var cloudSpaceUnpublishCmd = &cobra.Command{
	Use:   "unpublish",
	Args:  cobra.ExactArgs(1),
	Short: "Remove a previously published space from the cloud",
	Run:   func(cmd *cobra.Command, args []string) { ctrl.CloudSpaceUnpublish(args[0]) },
}

func init() {
	cloudCmd.AddCommand(cloudSpaceInfoCmd)
	cloudCmd.AddCommand(cloudSpacePublishCmd)
	cloudCmd.AddCommand(cloudSpaceUnpublishCmd)

	cloudSpacePublishCmd.Flags().StringVarP(&controllers.OrganizationOption, "organization", "o", "", "Publish under this organization")
}
