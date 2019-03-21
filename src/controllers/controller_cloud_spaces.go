package controllers

import (
	"fmt"
	"log"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools/console"
	"github.com/dplabs/cbox/src/tools/tty"
)

func (ctrl *CLIController) CloudSpaceInfo(spcSelectorStr string) {
	console.PrintAction("Retrieving info of an space")

	selector, err := models.ParseSelectorForCloud(spcSelectorStr)
	if err != nil {
		log.Fatalf("cloud: space info: %v", err)
	}

	space, err := ctrl.cloud.SpaceFind(selector)
	if err != nil {
		log.Fatalf("cloud: space info: %v", err)
	}

	console.PrintSpace(selector.String(), space)
}

func (ctrl *CLIController) CloudSpacePublish(spcSelectorStr string) {
	console.PrintAction("Publishing an space")

	selector, err := models.ParseSelectorMandatorySpace(spcSelectorStr)
	if err != nil {
		log.Fatalf("cloud: publish space: %v", err)
	}

	space, err := ctrl.findSpace(selector)
	if err != nil {
		log.Fatalf("cloud: publish space: %v", err)
	}

	if space.Selector.Namespace == "" {
		space.Selector.NamespaceType = models.TypeUser
		space.Selector.Namespace = ctrl.cloud.Login
	}

	previousSpace := space.Selector.Namespace

	if OrganizationOption != "" {
		space.Selector.NamespaceType = models.TypeOrganization
		space.Selector.Namespace = OrganizationOption
	}

	console.PrintSpace("Space to publish", space)

	if selector.Item != "" {
		commands := space.CommandList(selector.Item)
		if len(commands) == 0 {
			log.Fatalf("cloud: publish space: no local commands matched selector: %s", selector)
		}

		space.Entries = commands
	}

	// console.PrintCommandList("Containing these commands", space.Entries, false, false)

	if OrganizationOption != "" && previousSpace != OrganizationOption {
		console.PrintWarning(fmt.Sprintf("You're about to publish workspace '%s' under a different organization '%s'\n", space.String(), OrganizationOption))
	}

	if SkipQuestionsFlag || tty.Confirm("Publish?") {
		tty.Print("Publishing space '%s'...\n\n", space.String())

		err = ctrl.cloud.SpacePublish(space)
		if err != nil {
			log.Fatalf("cloud: publish space: %v", err)
		}

		ctrl.cleanOldSpaceFile(space, selector)

		core.Save(ctrl.cbox) // to store space's new namespace

		console.PrintSuccess("Space published successfully!")
	} else {
		console.PrintError("Publishing cancelled")
	}
}

func (ctrl *CLIController) CloudSpaceUnpublish(spcSelectorStr string) {
	console.PrintAction("Unpublishing an space")

	selector, err := models.ParseSelectorForCloud(spcSelectorStr)
	if err != nil {
		log.Fatalf("cloud: unpublish space: %v", err)
	}

	console.PrintSelector("Space to unpublish", selector)

	_, err = ctrl.findSpace(selector)
	if err == nil {
		console.PrintInfo("Local copy won't be deleted")
	} else {
		console.PrintWarning("You don't have a local copy of the space\n")
	}

	if SkipQuestionsFlag || tty.Confirm("Unpublish?") {
		tty.Print("Unpublishing space '%s'...\n\n", selector.String())

		err = ctrl.cloud.SpaceUnpublish(selector)
		if err != nil {
			log.Fatalf("cloud: unpublish space: %v", err)
		}

		console.PrintSuccess("Space unpublished successfully!")
	} else {
		console.PrintError("Unpublishing cancelled")
	}
}
