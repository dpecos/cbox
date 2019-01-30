package cli

import (
	"github.com/dplabs/cbox/src/controllers"
	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/cobra"
)

var cloudSpaceCmd = &cobra.Command{
	Use:     "spaces",
	Aliases: []string{"space"},
	Args:    cobra.ExactArgs(0),
	Short:   "Cloud operations for spaces",
	Long:    tools.Logo,
}

var cloudSpaceInfoCmd = &cobra.Command{
	Use:   "info",
	Args:  cobra.ExactArgs(1),
	Short: "Retrieve cloud info for an space",
	Run:   func(cmd *cobra.Command, args []string) { ctrl.CloudSpaceInfo(args) },
}

var cloudSpacePublishCmd = &cobra.Command{
	Use:     "publish",
	Aliases: []string{"push"},
	Args:    cobra.ExactArgs(1),
	Short:   "Publish a command, tag or a whole space",
	Run:     func(cmd *cobra.Command, args []string) { ctrl.CloudSpacePublish(args) },
}

var cloudSpaceUnpublishCmd = &cobra.Command{
	Use:   "unpublish",
	Args:  cobra.ExactArgs(1),
	Short: "Unpublish an space",
	Run:   func(cmd *cobra.Command, args []string) { ctrl.CloudSpaceUnpublish(args) },
}

var cloudSpacePullCmd = &cobra.Command{
	Use:   "pull",
	Args:  cobra.ExactArgs(1),
	Short: "Pull latest changes of a cloud space",
	Run:   func(cmd *cobra.Command, args []string) { ctrl.CloudSpacePull(args) },
}

var cloudSpaceCloneCmd = &cobra.Command{
	Use:   "clone",
	Args:  cobra.ExactArgs(1),
	Short: "Clone an space locally",
	Run:   func(cmd *cobra.Command, args []string) { ctrl.CloudSpaceClone(args) },
}

func init() {
	cloudCmd.AddCommand(cloudSpaceCmd)
	cloudSpaceCmd.AddCommand(cloudSpaceInfoCmd)
	cloudSpaceCmd.AddCommand(cloudSpacePublishCmd)
	cloudSpaceCmd.AddCommand(cloudSpaceUnpublishCmd)
	cloudSpaceCmd.AddCommand(cloudSpacePullCmd)
	cloudSpaceCmd.AddCommand(cloudSpaceCloneCmd)

	cloudSpacePublishCmd.Flags().StringVarP(&controllers.OrganizationOption, "organization", "o", "", "Publish under this organization")
}
