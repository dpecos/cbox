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

func (ctrl *CLIController) SpacesList(cmd *cobra.Command, args []string) {

	spaces := core.SpaceList()
	for _, space := range spaces {
		tools.PrintSpace(space)
	}
}

func (ctrl *CLIController) SpacesAdd(cmd *cobra.Command, args []string) {

	cbox := core.LoadCbox()

	space := tools.ConsoleReadSpace()
	cbox.SpaceCreate(space)

	core.PersistCbox(cbox)

	fmt.Println(aurora.Green("\nSpace successfully created!\n"))
	tools.PrintSpace(space)
}

func (ctrl *CLIController) SpacesEdit(cmd *cobra.Command, args []string) {

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

}

func (ctrl *CLIController) SpacesDelete(cmd *cobra.Command, args []string) {

	selector, err := models.ParseSelectorMandatorySpace(args[0])
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

}
