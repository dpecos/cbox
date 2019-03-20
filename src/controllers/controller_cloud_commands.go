package controllers

import (
	"fmt"
	"log"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools/console"
	"github.com/dplabs/cbox/src/tools/tty"
)

func (ctrl *CLIController) CloudCommandList(selectorStr string) {
	selector, err := models.ParseSelectorForCloud(selectorStr)
	if err != nil {
		log.Fatalf("cloud: list commands: invalid ctrl.cloud selector: %v", err)
	}

	commands, err := ctrl.cloud.CommandList(selector)
	if err != nil {
		log.Fatalf("cloud: list commands: %v", err)
	}

	console.PrintCommandList(selector.String(), commands, ListingsModeOption, ListingsSortOption)
}

func (ctrl *CLIController) CloudCommandCopy(cmdSelectorStr string, spcSelectorStr string) {
	console.PrintAction("Copying ctrl.cloud commands")

	cmdSelector, err := models.ParseSelectorForCloudCommand(cmdSelectorStr)
	if err != nil {
		log.Fatalf("cloud: copy command: invalid ctrl.cloud selector: %v", err)
	}

	spaceSelector, err := models.ParseSelectorMandatorySpace(spcSelectorStr)
	if err != nil {
		log.Fatalf("cloud: copy command: invalid space selector: %v", err)
	}

	space, err := ctrl.findSpace(spaceSelector)
	if err != nil {
		log.Fatalf("cloud: copy command: local space: %v", err)
	}

	commands, err := ctrl.cloud.CommandList(cmdSelector)
	if err != nil {
		log.Fatalf("cloud: copy command: retrieving matches: %v", err)
	}

	if len(commands) == 0 {
		console.PrintError(fmt.Sprintf("Command '%s' not found", cmdSelector))
	}

	console.PrintCommandList("Commands to copy", commands, "static", ListingsSortOption)

	if SkipQuestionsFlag || tty.Confirm(fmt.Sprintf("Copy these commands into %s?", spaceSelector)) {

		failures := false
		for _, command := range commands {
			err = space.CommandAdd(command)
			if err != nil {
				failures = true
				log.Printf("cloud: copy command: %v", err)
			}
		}

		core.Save(ctrl.cbox)

		if failures {
			console.PrintError("Some commands could not be stored")
		} else {
			console.PrintSuccess("Commands copied successfully!")
		}
	} else {
		console.PrintError("Copy cancelled")
	}
}
