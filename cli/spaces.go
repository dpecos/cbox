package cli

import (
	"github.com/dpecos/cmdbox/db"
	"github.com/dpecos/cmdbox/tools"
	"github.com/spf13/cobra"
)

var spacesCmd = &cobra.Command{
	Use:   "spaces",
	Short: "Show available spaces in your cmdbox",
	Run: func(cmd *cobra.Command, args []string) {
		cmdboxDB := db.Load(dbPath)
		defer cmdboxDB.Close()

		spaces := db.SpacesList()
		for _, space := range spaces {
			tools.PrintSpace(space)
		}
	},
}

func init() {
	rootCmd.AddCommand(spacesCmd)
}
