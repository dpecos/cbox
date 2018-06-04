package cli

import (
	"github.com/spf13/cobra"
)

var spacesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a space from your cmdbox",
	Run: func(cmd *cobra.Command, args []string) {

		// cmdboxDB := db.Load(dbPath)
		// defer cmdboxDB.Close()

		// id := tools.StringToUUID(args[0])
		// space := db.SpacesFind(id)

		// tools.PrintSpace(space)
		// if tools.Confirm(aurora.Red("Are you sure you want to delete this space and all its commands?").String()) {
		// 	db.SpacesDelete(id)
		// 	fmt.Println(aurora.Green("\nSpace deleted successfully!\n"))
		// }
	},
}

func init() {
	spacesCmd.AddCommand(spacesDeleteCmd)
}
