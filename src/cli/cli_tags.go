package cli

import (
	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/cobra"
)

var tagsCmd = &cobra.Command{
	Use:     "tags",
	Aliases: []string{"tag"},
	Args:    cobra.MaximumNArgs(1),
	Short:   "List all the tags in use within your cbox",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.TagsList(optionalSelector(args, 0)) },
}

var tagsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Args:  cobra.ExactArgs(1),
	Short: "Remove an specific tag from an space",
	Long:  tools.Logo,
	Run:   func(cmd *cobra.Command, args []string) { ctrl.TagsDelete(args[0]) },
}

func init() {
	rootCmd.AddCommand(tagsCmd)
	tagsCmd.AddCommand(tagsDeleteCmd)
}
