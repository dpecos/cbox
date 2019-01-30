package controllers

import (
	"log"
	"strings"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools"
	"github.com/dplabs/cbox/src/tools/console"
)

func (ctrl *CLIController) SpacesList(args []string) {
	for _, space := range ctrl.cbox.Spaces {
		tools.PrintSpace("", space)
	}
}

func (ctrl *CLIController) SpacesCreate(args []string) {
	console.PrintAction("Creating new space")

	space := tools.ConsoleReadSpace()

	err := ctrl.cbox.SpaceCreate(space)
	for err != nil {
		console.PrintError("Space already found in your cbox. Try a different one")
		space.Label = strings.ToLower(console.ReadString("Label", console.NOT_EMPTY_VALUES, console.ONLY_VALID_CHARS))
		space.Selector.Space = space.Label
		err = ctrl.cbox.SpaceCreate(space)
	}

	core.Save(ctrl.cbox)

	tools.PrintSpace("New space", space)

	console.PrintSuccess("Space successfully created!")
}

func (ctrl *CLIController) SpacesEdit(args []string) {
	console.PrintAction("Editing an space")

	selector, err := models.ParseSelectorMandatorySpace(args[0])
	if err != nil {
		log.Fatalf("edit space: %v", err)
	}

	space, err := ctrl.findSpace(selector)
	if err != nil {
		log.Fatalf("edit space: %v", err)
	}

	tools.PrintSpace("Space to edit", space)

	tools.ConsoleEditSpace(space)

	tools.PrintSpace("Space after edition", space)

	if SkipQuestionsFlag || console.Confirm("Update?") {

		err := ctrl.cbox.SpaceEdit(space, selector.Namespace, selector.Space)
		for err != nil {
			console.PrintError("Label already found in your cbox. Try a different one")
			space.Label = strings.ToLower(console.ReadString("Label", console.NOT_EMPTY_VALUES, console.ONLY_VALID_CHARS))
			err = ctrl.cbox.SpaceEdit(space, selector.Namespace, selector.Space)
		}

		ctrl.cleanOldSpaceFile(space, selector)

		core.Save(ctrl.cbox)

		console.PrintSuccess("Space updated successfully!")
	} else {
		console.PrintError("Edition cancelled")
	}
}

func (ctrl *CLIController) SpacesDestroy(args []string) {
	console.PrintAction("Destroying an space")

	selector, err := models.ParseSelectorMandatorySpace(args[0])
	if err != nil {
		log.Fatalf("destroy space: %v", err)
	}

	space, err := ctrl.findSpace(selector)
	if err != nil {
		log.Fatalf("destroy space: %v", err)
	}

	tools.PrintSpace("Space to destroy", space)

	if SkipQuestionsFlag || console.Confirm("Are you sure you want to destroy this space?") {
		err = ctrl.cbox.SpaceDestroy(space)
		if err != nil {
			log.Fatalf("destroy space: %v", err)
		}
		core.DeleteSpaceFile(space.Selector)

		console.PrintSuccess("Space destroyed successfully!")
	} else {
		console.PrintError("Deletion cancelled")
	}
}