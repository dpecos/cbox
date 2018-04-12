package cli

import (
	"github.com/dpecos/cmdbox/db"
	"github.com/dpecos/cmdbox/tools"
	"github.com/spf13/cobra"
)

var (
	viewSnippet bool
	filterTag   string
)
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List the content of your cmdbox",
	Run: func(cmd *cobra.Command, args []string) {
		cmdboxDB := db.Load(dbPath)
		defer cmdboxDB.Close()

		cmds := db.List(filterTag)
		for _, cmd := range cmds {
			tools.PrintCommand(cmd, viewSnippet, false)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&viewSnippet, "view", "v", false, "Show all details about commands")
	listCmd.Flags().StringVarP(&filterTag, "tag", "t", "", "List commands only of specified tag")
}
