package cli

import (
	"fmt"

	"github.com/dpecos/cmdbox/db"
	"github.com/dpecos/cmdbox/tools"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Args:    cobra.ExactArgs(1),
	Short:   "Delete a command from your cmdbox",
	Run: func(cmd *cobra.Command, args []string) {
		cmdboxDB := db.Load(dbPath)
		defer cmdboxDB.Close()

		id := tools.StringToInt(args[0])
		command := db.Find(id)

		tools.PrintCommand(command, true, false)
		if tools.Confirm(aurora.Red("Are you sure you want to delete this command?").String()) {
			db.Delete(id)
			fmt.Println(aurora.Green("\nCommand successfully deleted!\n"))
		}

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
