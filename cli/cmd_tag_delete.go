package cli

import (
	"github.com/dpecos/cmdbox/db"
	"github.com/dpecos/cmdbox/tools"
	"github.com/spf13/cobra"
)

var tagDeleteCmd = &cobra.Command{
	Use:   "delete",
	Args:  cobra.MinimumNArgs(2),
	Short: "Delete tags from a command",
	Run: func(cmd *cobra.Command, args []string) {
		cmdboxDB := db.Load(dbPath)
		defer cmdboxDB.Close()

		cmdID := tools.StringToInt(args[0])
		tag := args[1]

		db.UnassignTag(cmdID, tag)

		command := db.Find(cmdID)
		tools.PrintCommand(command, false, false)
	},
}

func init() {
	tagCmd.AddCommand(tagDeleteCmd)
}
