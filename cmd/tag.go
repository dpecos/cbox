package cmd

import (
	"github.com/dpecos/cmdbox/db"
	"github.com/dpecos/cmdbox/tools"
	"github.com/spf13/cobra"
)

var deleteTag string
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Add tags to a command",
	Run: func(cmd *cobra.Command, args []string) {
		cmdboxDB := db.Load(dbPath)
		defer cmdboxDB.Close()

		cmdID := tools.StringToInt(args[0])
		if deleteTag != "" {
			db.UnassignTag(cmdID, deleteTag)
		} else {
			for _, tag := range args[1:] {
				db.AssignTag(cmdID, tag)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)

	tagCmd.Flags().StringVarP(&deleteTag, "delete", "d", "", "Remove specified tag from command")
}
