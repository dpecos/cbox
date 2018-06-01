package cli

import (
	"github.com/dpecos/cmdbox/db"
	"github.com/dpecos/cmdbox/tools"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:     "tag",
	Aliases: []string{"t"},
	Args:    cobra.MinimumNArgs(2),
	Short:   "Add tags to a command",
	Run: func(cmd *cobra.Command, args []string) {
		cmdboxDB := db.Load(dbPath)
		defer cmdboxDB.Close()

		cmdID := tools.StringToInt(args[0])

		for _, tag := range args[1:] {
			db.AssignTag(cmdID, tag)
		}

		command := db.Find(cmdID)
		tools.PrintCommand(command, false, false)
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
}
