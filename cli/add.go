package cli

import (
	"strings"

	"github.com/dpecos/cmdbox/db"
	"github.com/dpecos/cmdbox/models"
	"github.com/dpecos/cmdbox/tools"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.ExactArgs(0),
	Short: "Add a new command to your cmdbox",
	Run: func(cmd *cobra.Command, args []string) {

		cmdboxDB := db.Load(dbPath)
		defer cmdboxDB.Close()

		command := models.Cmd{
			Title:       tools.ReadString("Title"),
			Description: tools.ReadStringMulti("Description"),
			URL:         tools.ReadString("URL"),
			Cmd:         tools.ReadStringMulti("Command / snippet"),
		}

		id := db.Add(command)

		tags := tools.ReadString("Tags (separated by space)")
		for _, tag := range strings.Split(tags, " ") {
			db.AssignTag(id, tag)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
