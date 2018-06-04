package cli

import (
	"github.com/spf13/cobra"
)

var spacesAddCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.ExactArgs(0),
	Short: "Add a new space to your cmdbox",
	Run: func(cmd *cobra.Command, args []string) {

		// cmdboxDB := db.Load(dbPath)
		// defer cmdboxDB.Close()

		// space := models.Space{
		// 	Name:  tools.ReadString("Name"),
		// 	Title: tools.ReadString("Title"),
		// }

		// id := db.SpacesCreate(space)

		// fmt.Println(aurora.Green("\nSpace created successfully!\n"))
		// space = db.SpacesFind(id)
		// tools.PrintSpace(space)
	},
}

func init() {
	spacesCmd.AddCommand(spacesAddCmd)
}
