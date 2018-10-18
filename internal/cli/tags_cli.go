package cli

import (
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/cobra"
)

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Args:  cobra.MaximumNArgs(1),
	Short: "List the tags available in your cbox",
	Long:  tools.Logo,
	Run:   ctrl.TagsList,
}

var tagsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Args:  cobra.ExactArgs(1),
	Short: "Delete specified tag",
	Long:  tools.Logo,
	Run:   ctrl.TagsDelete,
}

func init() {
	rootCmd.AddCommand(tagsCmd)
	tagsCmd.AddCommand(tagsDeleteCmd)
}