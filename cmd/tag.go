package cmd

import (
	"log"
	"strconv"

	"github.com/dpecos/cmdbox/db"
	"github.com/spf13/cobra"
)

var deleteTag string
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Add tags to a command",
	Run: func(cmd *cobra.Command, args []string) {
		cmdboxDB := db.Load(dbPath)
		defer cmdboxDB.Close()

		if deleteTag != "" {
			cmdID := readIntParam(args, 0)
			db.UnassignTag(cmdID, deleteTag)
		} else {
			cmdID := readIntParam(args, 0)
			for _, tag := range args[1:] {
				db.AssignTag(cmdID, tag)
			}
		}
	},
}

func readIntParam(args []string, pos int) int64 {
	param, err := strconv.ParseInt(args[pos], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return param
}

func init() {
	rootCmd.AddCommand(tagCmd)

	tagCmd.Flags().StringVarP(&deleteTag, "delete", "d", "", "Remove specified tag from command")
}
