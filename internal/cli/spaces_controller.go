package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/dpecos/cbox/internal/core"
	"github.com/dpecos/cbox/pkg/models"
	"github.com/dpecos/cbox/tools"
	"github.com/dpecos/cbox/tools/console"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) SpacesList(cmd *cobra.Command, args []string) {

	spaces := core.SpaceListFiles()
	for _, space := range spaces {
		tools.PrintSpace("", space)
	}
}

func (ctrl *CLIController) SpacesCreate(cmd *cobra.Command, args []string) {

	cbox := core.LoadCbox("")

	fmt.Println("Creating new space")
	space := tools.ConsoleReadSpace()

	err := cbox.SpaceCreate(space)
	for err != nil {
		console.PrintError("Space already found in your cbox. Try a different one")
		space.Label = strings.ToLower(console.ReadString("Label"))
		err = cbox.SpaceCreate(space)
	}

	core.PersistCbox(cbox)

	tools.PrintSpace("New space", space)

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

	tools.PrintSpace("Space to edit", space)

	tools.ConsoleEditSpace(space)

	tools.PrintSpace("Space after edition", space)

	if console.Confirm("Update?") {

		err := cbox.SpaceEdit(space, selector.Space)
		for err != nil {
			console.PrintError("Label already found in your cbox. Try a different one")
			space.Label = strings.ToLower(console.ReadString("Label"))
			err = cbox.SpaceEdit(space, selector.Space)
		}

		if space.Label != selector.Space {
			spaceToDelete := &models.Space{
				Label: selector.Space,
			}
			core.SpaceDeleteFile(spaceToDelete)
		}

		core.PersistCbox(cbox)
		console.PrintSuccess("Space updated successfully!")
	} else {
		console.PrintError("Edition cancelled")
	}
}

func (ctrl *CLIController) SpacesDestroy(cmd *cobra.Command, args []string) {

	selector, err := models.ParseSelectorMandatorySpace(args[0])
	if err != nil {
		log.Fatalf("destroy space: %v", err)
	}

	cbox := core.LoadCbox("")

	space, err := cbox.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("destroy space: %v", err)
	}

	tools.PrintSpace("Space to destroy", space)

	if console.Confirm("Are you sure you want to destroy this space?") {
		// fix issue #52: when a space is removed, pointers to that position of memory change values
		s := models.Space{
			Label: space.Label,
		}
		s.ID = space.ID

		err = cbox.SpaceDestroy(&s)
		if err != nil {
			log.Fatalf("destroy space: %v", err)
		}
		core.SpaceDeleteFile(&s)

		console.PrintSuccess("Space destroyed successfully!")
	} else {
		console.PrintError("Deletion cancelled")
	}
}
