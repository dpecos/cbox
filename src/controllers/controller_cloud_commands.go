package controllers

import (
	"fmt"
	"log"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools"
	"github.com/dplabs/cbox/src/tools/console"
)

func (ctrl *CLIController) CloudCommandList(args []string) {
	selector, err := models.ParseSelectorForCloud(args[0])
	if err != nil {
		log.Fatalf("ctrl.cloud: list commands: invalid ctrl.cloud selector: %v", err)
	}

	commands, err := ctrl.cloud.CommandList(selector)
	if err != nil {
		log.Fatalf("ctrl.cloud: list commands: %v", err)
	}

	tools.PrintCommandList(selector.String(), commands, ShowCommandsSourceFlag, false)
}

func (ctrl *CLIController) CloudCommandCopy(args []string) {
	console.PrintAction("Copying ctrl.cloud commands")

	cmdSelector, err := models.ParseSelectorForCloudCommand(args[0])
	if err != nil {
		log.Fatalf("ctrl.cloud: copy command: invalid ctrl.cloud selector: %v", err)
	}

	spaceSelector, err := models.ParseSelectorMandatorySpace(args[1])
	if err != nil {
		log.Fatalf("ctrl.cloud: copy command: invalid space selector: %v", err)
	}

	space, err := ctrl.findSpace(spaceSelector)
	if err != nil {
		log.Fatalf("ctrl.cloud: copy command: local space: %v", err)
	}

	commands, err := ctrl.cloud.CommandList(cmdSelector)
	if err != nil {
		log.Fatalf("ctrl.cloud: copy command: retrieving matches: %v", err)
	}

	if len(commands) == 0 {
		console.PrintError(fmt.Sprintf("Command '%s' not found", cmdSelector))
	}

	tools.PrintCommandList("Commands to copy", commands, false, false)

	if SkipQuestionsFlag || console.Confirm(fmt.Sprintf("Copy these commands into %s?", spaceSelector)) {

		failures := false
		for _, command := range commands {
			err = space.CommandAdd(command)
			if err != nil {
				failures = true
				log.Printf("ctrl.cloud: copy command: %v", err)
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
