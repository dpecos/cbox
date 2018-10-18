package cli

import (
	"github.com/dpecos/cbox/internal/pkg"
	"github.com/spf13/cobra"
)

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Args:  cobra.MaximumNArgs(1),
	Short: "List the tags available in your cbox",
	Long:  pkg.Logo,
	Run:   ctrl.TagsList,
}

var tagsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Args:  cobra.ExactArgs(1),
	Short: "Delete specified tag",
	Long:  pkg.Logo,
	Run:   ctrl.TagsDelete,
}

func init() {
	rootCmd.AddCommand(tagsCmd)
	tagsCmd.AddCommand(tagsDeleteCmd)
}
