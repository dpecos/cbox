package cmd

import (
	"fmt"
	"strings"

	"github.com/dpecos/cmdbox/db"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var viewSnippet bool
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List the content of your cmdbox",
	Run: func(cmd *cobra.Command, args []string) {
		cmdboxDB := db.Load(dbPath)
		defer cmdboxDB.Close()

		cmds := db.List()
		for _, cmd := range cmds {
			if len(cmd.Tags) != 0 {
				tags := strings.Join(cmd.Tags, ", ")
				fmt.Printf("%d - (%s) %s - %s\n", aurora.Red(aurora.Bold(cmd.ID)), aurora.Brown(tags), aurora.Blue(aurora.Bold(cmd.Title)), aurora.Green(cmd.CreatedAt))
			} else {
				fmt.Printf("%d - %s - %s\n", aurora.Red(aurora.Bold(cmd.ID)), aurora.Blue(aurora.Bold(cmd.Title)), aurora.Green(cmd.CreatedAt))

			}
			if viewSnippet {
				if cmd.Description != "" {
					fmt.Printf("\n%s\n", aurora.Cyan(cmd.Description))
				}
				if cmd.URL != "" {
					fmt.Printf("\n%s\n", aurora.Blue(cmd.URL))
				}
				fmt.Printf("\n%s\n\n", cmd.Cmd)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&viewSnippet, "view", "v", false, "Show code snippet")
}
