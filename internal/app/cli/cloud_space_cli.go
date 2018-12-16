package cli

import (
	"github.com/dplabs/cbox/internal/pkg"
	"github.com/spf13/cobra"
)

var cloudSpaceCmd = &cobra.Command{
	Use:     "spaces",
	Aliases: []string{"space"},
	Args:    cobra.ExactArgs(0),
	Short:   "Cloud operations for spaces",
	Long:    pkg.Logo,
}

var cloudSpaceInfoCmd = &cobra.Command{
	Use:   "info",
	Args:  cobra.ExactArgs(1),
	Short: "Retrieve cloud info for an space",
	Run:   ctrl.CloudSpaceInfo,
}

var cloudSpacePublishCmd = &cobra.Command{
	Use:     "publish",
	Aliases: []string{"push"},
	Args:    cobra.ExactArgs(1),
	Short:   "Publish a command, tag or a whole space",
	Run:     ctrl.CloudSpacePublish,
}

var cloudSpaceUnpublishCmd = &cobra.Command{
	Use:   "unpublish",
	Args:  cobra.ExactArgs(1),
	Short: "Unpublish an space",
	Run:   ctrl.CloudSpaceUnpublish,
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

func init() {
	cloudCmd.AddCommand(cloudSpaceCmd)
	cloudSpaceCmd.AddCommand(cloudSpaceInfoCmd)
	cloudSpaceCmd.AddCommand(cloudSpacePublishCmd)
	cloudSpaceCmd.AddCommand(cloudSpaceUnpublishCmd)
	cloudSpaceCmd.AddCommand(cloudSpacePullCmd)
	cloudSpaceCmd.AddCommand(cloudSpaceCloneCmd)
}
