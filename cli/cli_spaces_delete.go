package cli

import (
	"fmt"
	"log"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools"
	"github.com/dpecos/cbox/tools/console"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var spacesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a space from your cbox",
	Long:  tools.Logo,
	Run: func(cmd *cobra.Command, args []string) {

		selector, err := models.ParseSelector(args[0])
		if err != nil {
			log.Fatal("Could not parse selector", err)
		}

		cbox := core.LoadCbox()

		space := cbox.SpaceFind(selector.Space)

		tools.PrintSpace(space)
		if console.Confirm(aurora.Red("Are you sure you want to delete this space?").String()) {
			core.SpaceDelete(space)
			fmt.Println(aurora.Green("\nSpace successfully deleted!\n"))
		}

	},
}

func init() {
	spacesCmd.AddCommand(spacesDeleteCmd)
}
