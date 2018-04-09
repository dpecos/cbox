package cmd

import (
	"log"
	"strconv"

	"github.com/dpecos/cmdbox/db"
	"github.com/spf13/cobra"
)

// tagCmd represents the tag command
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Add / Remove tags to a command",
	Run: func(cmd *cobra.Command, args []string) {
		cmdboxDB := db.Load(dbPath)
		defer cmdboxDB.Close()

		cmdID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		db.AssignTag(cmdID, args[1])
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
}
