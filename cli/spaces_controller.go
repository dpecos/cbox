package cli

import (
	"fmt"
	"log"
	"strings"

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

	cbox := core.LoadCbox("")

	fmt.Println("Creating new space")
	space := tools.ConsoleReadSpace()

	err := cbox.SpaceAdd(space)
	for err != nil {
		console.PrintError("Space already found in your cbox. Try a different one")
		space.Label = strings.ToLower(console.ReadString("Label"))
		err = cbox.SpaceAdd(space)
	}

	core.PersistCbox(cbox)

	fmt.Printf("\n--- New space ---\n")
	tools.PrintSpace(space)
	fmt.Printf("-----\n\n")

	console.PrintSuccess("Space successfully created!")
}

func (ctrl *CLIController) SpacesEdit(cmd *cobra.Command, args []string) {

	selector, err := models.ParseSelectorMandatorySpace(args[0])
	if err != nil {
		log.Fatalf("edit space: %v", err)
	}

	cbox := core.LoadCbox("")

	space, err := cbox.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("edit space: %v", err)
	}

	fmt.Printf("--- Space to edit ---\n")
	tools.PrintSpace(space)
	fmt.Printf("-----\n\n")

	tools.ConsoleEditSpace(space)

	if space.Label != selector.Space {
		for len(cbox.SpaceLabels()) != len(cbox.Spaces) {
			console.PrintError("Label already found in your cbox. Try a different one")
			space.Label = strings.ToLower(console.ReadString("Label"))
		}
	}

	fmt.Printf("--- Space after edited values ---\n")
	tools.PrintSpace(space)
	fmt.Printf("-----\n\n")

	if console.Confirm("Update?") {
		spaceToDelete := &models.Space{
			Label: selector.Space,
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
		log.Fatalf("delete space: %v", err)
	}

	cbox := core.LoadCbox("")

	space, err := cbox.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("delete space: %v", err)
	}

	fmt.Printf("\n--- Space to delete ---\n")
	tools.PrintSpace(space)
	fmt.Printf("-----\n\n")

	if console.Confirm("Are you sure you want to delete this space?") {
		core.SpaceDelete(space)
		console.PrintSuccess("Space deleted successfully!")
	} else {
		console.PrintError("Deletion cancelled")
	}
}
