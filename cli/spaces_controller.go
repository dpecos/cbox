package cli

import (
	"fmt"
	"log"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools"
	"github.com/dpecos/cbox/tools/console"
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

	fmt.Println("Creating new space")
	space := tools.ConsoleReadSpace()

	cbox.SpaceCreate(space)
	core.PersistCbox(cbox)

	fmt.Println("\n--- New space ---")
	tools.PrintSpace(space)
	fmt.Println("-----\n")

	console.PrintSuccess("Space successfully created!")
}

func (ctrl *CLIController) SpacesEdit(cmd *cobra.Command, args []string) {

	selector, err := models.ParseSelectorMandatorySpace(args[0])
	if err != nil {
		log.Fatalf("Could not parse selector: %s", err)
	}

	cbox := core.LoadCbox()

	space := cbox.SpaceFind(selector.Space)

	fmt.Printf("Editing space with Name %s\n", space.Name)
	tools.ConsoleEditSpace(space)

	cbox.SpaceEdit(space, selector.Space)

	fmt.Println("--- Space after edited values ---")
	tools.PrintSpace(space)
	fmt.Println("-----\n")

	if console.Confirm("Update?") {
		spaceToDelete := &models.Space{
			Name: selector.Space,
		}
		core.SpaceDelete(spaceToDelete)

		core.PersistCbox(cbox)
		console.PrintSuccess("Space updated successfully!")
	} else {
		console.PrintError("Edition cancelled")
	}
}

func (ctrl *CLIController) SpacesDelete(cmd *cobra.Command, args []string) {

	selector, err := models.ParseSelectorMandatorySpace(args[0])
	if err != nil {
		log.Fatal("Could not parse selector", err)
	}

	cbox := core.LoadCbox()

	space := cbox.SpaceFind(selector.Space)

	fmt.Println("\n--- Space to delete ---")
	tools.PrintSpace(space)
	fmt.Println("-----\n")

	if console.Confirm("Are you sure you want to delete this space?") {
		core.SpaceDelete(space)
		console.PrintSuccess("Space deleted successfully!")
	} else {
		console.PrintError("Deletion cancelled")
	}
}
