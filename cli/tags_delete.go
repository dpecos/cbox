package cli

import (
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/cobra"
)

var tagsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete specified tag",
	Args:  cobra.ExactArgs(1),
	Long:  tools.Logo,
	Run: func(cmd *cobra.Command, args []string) {
		// cmdboxDB := db.Load(dbPath)
		// defer cmdboxDB.Close()

		// tag := args[0]

		// db.TagsDelete(tag)
	},
}

func init() {
	tagsCmd.AddCommand(tagsDeleteCmd)
}
