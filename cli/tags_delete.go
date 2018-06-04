package cli

import (
	"github.com/spf13/cobra"
)

var tagsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete specified tag",
	Args:  cobra.ExactArgs(1),
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
