package cli

import (
	"github.com/spf13/cobra"
)

var spacesEditCmd = &cobra.Command{
	Use:   "edit",
	Args:  cobra.ExactArgs(1),
	Short: "Edit an space from your cmdbox",
	Run: func(cmd *cobra.Command, args []string) {
		// cmdboxDB := db.Load(dbPath)
		// defer cmdboxDB.Close()

		// id := tools.StringToUUID(args[0])
		// space := db.SpacesFind(id)

		// space.Name = tools.EditString("Name", space.Name)
		// space.Title = tools.EditString("Title", space.Title)

		// tools.PrintSpace(space)
		// if tools.Confirm("Update?") {
		// 	db.SpacesUpdate(space)
		// } else {
		// 	fmt.Println("Cancelled")
		// }

	},
}

func init() {
	spacesCmd.AddCommand(spacesEditCmd)
}
