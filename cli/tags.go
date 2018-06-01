package cli

import (
	"github.com/dpecos/cmdbox/db"
	"github.com/dpecos/cmdbox/tools"
	"github.com/spf13/cobra"
)

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "List the tags available in your cmdbox",
	Run: func(cmd *cobra.Command, args []string) {
		cmdboxDB := db.Load(dbPath)
		defer cmdboxDB.Close()

		cmds := db.TagsList()
		for _, cmd := range cmds {
			tools.PrintTag(cmd)
		}
	},
}

func init() {
	rootCmd.AddCommand(tagsCmd)
}
