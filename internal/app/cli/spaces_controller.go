package cli

import (
	"log"
	"strings"

	"github.com/dplabs/cbox/internal/app/core"
	"github.com/dplabs/cbox/internal/pkg"
	"github.com/dplabs/cbox/internal/pkg/console"
	"github.com/dplabs/cbox/pkg/models"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) SpacesList(cmd *cobra.Command, args []string) {
	for _, space := range cboxInstance.Spaces {
		pkg.PrintSpace("", space)
	}
}

func (ctrl *CLIController) SpacesCreate(cmd *cobra.Command, args []string) {
	console.PrintAction("Creating new space")

	space := pkg.ConsoleReadSpace()

	err := cboxInstance.SpaceCreate(space)
	for err != nil {
		console.PrintError("Space already found in your cbox. Try a different one")
		space.Label = strings.ToLower(console.ReadString("Label", console.NOT_EMPTY_VALUES, console.ONLY_VALID_CHARS))
		err = cboxInstance.SpaceCreate(space)
	}

	core.Save(cboxInstance)

	pkg.PrintSpace("New space", space)

	console.PrintSuccess("Space successfully created!")
}

func (ctrl *CLIController) SpacesEdit(cmd *cobra.Command, args []string) {
	console.PrintAction("Editing an space")

	selector, err := models.ParseSelectorMandatorySpace(args[0])
	if err != nil {
		log.Fatalf("edit space: %v", err)
	}

	space, err := cboxInstance.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("edit space: %v", err)
	}

	pkg.PrintSpace("Space to edit", space)

	pkg.ConsoleEditSpace(space)

	pkg.PrintSpace("Space after edition", space)

	if console.Confirm("Update?") {

		err := cboxInstance.SpaceEdit(space, selector.Space)
		for err != nil {
			console.PrintError("Label already found in your cbox. Try a different one")
			space.Label = strings.ToLower(console.ReadString("Label", console.NOT_EMPTY_VALUES, console.ONLY_VALID_CHARS))
			err = cboxInstance.SpaceEdit(space, selector.Space)
		}

		if space.Label != selector.Space {
			spaceToDelete := models.Space{
				Label: selector.Space,
			}
			core.DeleteSpaceFile(&spaceToDelete)
		}

		core.Save(cboxInstance)
		console.PrintSuccess("Space updated successfully!")
	} else {
		console.PrintError("Edition cancelled")
	}
}

func (ctrl *CLIController) SpacesDestroy(cmd *cobra.Command, args []string) {
	console.PrintAction("Destroying an space")

	selector, err := models.ParseSelectorMandatorySpace(args[0])
	if err != nil {
		log.Fatalf("destroy space: %v", err)
	}

	space, err := cboxInstance.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("destroy space: %v", err)
	}

	pkg.PrintSpace("Space to destroy", space)

	if console.Confirm("Are you sure you want to destroy this space?") {
		// fix issue #52: when a space is removed, pointers to that position of memory change values
		s := models.Space{
			Label: space.Label,
		}
		s.ID = space.ID

		err = cboxInstance.SpaceDestroy(&s)
		if err != nil {
			log.Fatalf("destroy space: %v", err)
		}
		core.DeleteSpaceFile(&s)

		console.PrintSuccess("Space destroyed successfully!")
	} else {
		console.PrintError("Deletion cancelled")
	}
}
