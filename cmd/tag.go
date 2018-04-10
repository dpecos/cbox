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
			cmdID := readParamInt(args, 0)

			db.UnassignTag(cmdID, deleteTag)
		} else {
			cmdID := readParamInt(args, 1)
			tag := readParam(args, 0)

			db.AssignTag(cmdID, tag)
		}
	},
}

func readParam(args []string, pos int) string {
	param := args[pos]
	return param
}

func readParamInt(args []string, pos int) int {
	param, err := strconv.Atoi(readParam(args, pos))
	if err != nil {
		log.Fatal(err)
	}
	return param
}

func init() {
	rootCmd.AddCommand(tagCmd)

	tagCmd.Flags().StringVarP(&deleteTag, "delete", "d", "", "Remove specified tag from command")
}
