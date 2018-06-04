package cli

import (
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Args:  cobra.ExactArgs(1),
	Short: "Edit a command from your cbox",
	Run: func(cmd *cobra.Command, args []string) {
		// cmdboxDB := db.Load(dbPath)
		// defer cmdboxDB.Close()

		// cmdID := tools.StringToInt(args[0])
		// command := db.Find(cmdID)

		// command.Title = tools.EditString("Title", command.Title)
		// command.Description = tools.EditStringMulti("Description", command.Description)
		// command.URL = tools.EditString("URL", command.URL)
		// command.Cmd = tools.EditStringMulti("Command / Snippet", command.Cmd)

		// tools.PrintCommand(command, true, false)
		// if tools.Confirm("Update?") {
		// 	db.Update(command)
		// } else {
		// 	fmt.Println("Cancelled")
		// }

	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
