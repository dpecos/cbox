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

var spacesEditCmd = &cobra.Command{
	Use:   "edit",
	Args:  cobra.ExactArgs(1),
	Short: "Edit an space from your cbox",
	Long:  tools.Logo,
	Run: func(cmd *cobra.Command, args []string) {

		selector, err := models.ParseSelectorMandatorySpace(args[0])
		if err != nil {
			log.Fatalf("Could not parse selector: %s", err)
		}

		cbox := core.LoadCbox()

		space := cbox.SpaceFind(selector.Space)
		tools.ConsoleEditSpace(space)

		cbox.SpaceEdit(space, selector.Space)

		tools.PrintSpace(space)
		if console.Confirm("Update?") {
			spaceToDelete := &models.Space{
				Name: selector.Space,
			}
			core.SpaceDelete(spaceToDelete)

			core.PersistCbox(cbox)
			fmt.Println(aurora.Green("\nSpace updated successfully!\n"))
		} else {
			fmt.Println("Cancelled")
		}

	},
}

func init() {
	spacesCmd.AddCommand(spacesEditCmd)
}
