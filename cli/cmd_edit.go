package cli

import (
	"fmt"

	"github.com/dpecos/cmdbox/db"
	"github.com/dpecos/cmdbox/tools"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Args:  cobra.ExactArgs(1),
	Short: "Edit a command from your cmdbox",
	Run: func(cmd *cobra.Command, args []string) {
		cmdboxDB := db.Load(dbPath)
		defer cmdboxDB.Close()

		cmdID := tools.StringToInt(args[0])
		command := db.Find(cmdID)

		command.Title = tools.EditString("Title", command.Title)
		command.Description = tools.EditStringMulti("Description", command.Description)
		command.URL = tools.EditString("URL", command.URL)
		command.Cmd = tools.EditStringMulti("Command / Snippet", command.Cmd)

		tools.PrintCommand(command, true, false)
		if tools.Confirm("Update?") {
			db.Update(command)
		} else {
			fmt.Println("Cancelled")
		}

	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
