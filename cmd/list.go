package cmd

import (
	"fmt"

	"github.com/dpecos/cmdbox/db"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		cmdboxDB := db.Load(dbPath)
		defer cmdboxDB.Close()

		cmds := db.List()
		for _, cmd := range cmds {
			fmt.Printf("%d - %s - %s\n\n%s\n\n", aurora.Red(aurora.Bold(cmd.ID)), aurora.Blue(aurora.Bold(cmd.Title)), aurora.Green(cmd.CreatedAt), cmd.Cmd)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
