package cli

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"

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
			Cmd:         tools.ReadStringMulti("Command / Snippet"),
		}
		tags := tools.ReadString("Tags (separated by space)")

		id := db.Add(command)

		for _, tag := range strings.Split(tags, " ") {
			if tag != "" {
				db.AssignTag(id, tag)
			}
		}

		fmt.Println(aurora.Green("\nCommand stored successfully!\n"))
		command = db.Find(id)
		tools.PrintCommand(command, true, false)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
