package cli

import (
	"github.com/spf13/cobra"
)

var spacesCmd = &cobra.Command{
	Use:   "spaces",
	Short: "Show available spaces in your cbox",
	Run: func(cmd *cobra.Command, args []string) {
		// cmdboxDB := db.Load(dbPath)
		// defer cmdboxDB.Close()

		// spaces := db.SpacesList()
		// for _, space := range spaces {
		// 	tools.PrintSpace(space)
		// }
	},
}

func init() {
	rootCmd.AddCommand(spacesCmd)
}
