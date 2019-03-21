package controllers

import (
	"fmt"
	"log"
	"strings"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools/console"
	"github.com/dplabs/cbox/src/tools/tty"
)

func (ctrl *CLIController) CloudCommandList(selectorStr string) {
	selector, err := models.ParseSelectorForCloud(selectorStr)
	if err != nil {
		log.Fatalf("cloud: list commands: invalid cloud selector: %v", err)
	}

	commands, err := ctrl.cloud.CommandList(selector)
	if err != nil {
		log.Fatalf("cloud: list commands: %v", err)
	}

	if ListingsModeOption == "interactive" {
		ListingsModeOption = "interactive-remote"
	}
	console.PrintCommandList(selector.String(), commands, ListingsModeOption, ListingsSortOption)
}

func (ctrl *CLIController) cloneSpace(cloudSelector *models.Selector, commands []*models.Command) {
	console.PrintInfo(fmt.Sprintf("Cloning remote space '%s'...\n", cloudSelector.String()))

	space, err := ctrl.cloud.SpaceFind(cloudSelector)
	if err != nil {
		log.Fatalf("cloud: copy commands: %v", err)
	}
	space.Selector.Namespace = ""
	space.Selector.NamespaceType = models.TypeNone

	space.Entries = commands

	err = ctrl.cbox.SpaceCreate(space)
	for err != nil {
		console.PrintError("Space already found in your cbox. Try a different one")
		space.Label = strings.ToLower(console.ReadString("Label", console.NOT_EMPTY_VALUES, console.ONLY_VALID_CHARS))
		space.Selector.Space = space.Label
		space.ID = space.Selector.String()
		err = ctrl.cbox.SpaceCreate(space)
	}

	core.Save(ctrl.cbox)

	console.PrintSuccess(fmt.Sprintf("Space cloned successfully into '%s'!", space.Selector.String()))
}

func (ctrl *CLIController) copyCommands(spaceSelector *models.Selector, commands []*models.Command) {

	console.PrintInfo(fmt.Sprintf("Copying commands into existing space '%s'...\n", spaceSelector.String()))

	space, err := ctrl.findSpace(spaceSelector)
	if err != nil {
		log.Fatalf("cloud: copy command: %v", err)
	}

	failures := false
	for _, command := range commands {
		err = space.CommandAdd(command, ForceFlag)
		if err != nil {
			failures = true
			log.Printf("cloud: copy command: %v", err)
		}
	}

	core.Save(ctrl.cbox)

	if failures {
		console.PrintError("Some commands could not be stored")
	} else {
		console.PrintSuccess(fmt.Sprintf("Commands copied successfully into '%s'!", space.Selector.String()))
	}
}

func (ctrl *CLIController) CloudCopy(cloudSelectorStr string, spcSelectorStr *string) {
	console.PrintAction("Copying cloud commands")

	cloudSelector, err := models.ParseSelectorForCloud(cloudSelectorStr)
	if err != nil {
		log.Fatalf("cloud: copy command: invalid cloud selector: %v", err)
	}

	commands, err := ctrl.cloud.CommandList(cloudSelector)
	if err != nil {
		log.Fatalf("cloud: copy command: retrieving matches: %v", err)
	}

	if len(commands) == 0 {
		console.PrintError(fmt.Sprintf("Command '%s' not found", cloudSelector))
	}

	cloneRemoteSpace := spcSelectorStr == nil && cloudSelector.Item == ""

	var spaceSelector *models.Selector

	if cloneRemoteSpace {
		console.PrintWarning("Creating a new local workspace, clone of the remote.\n")
	} else {
		s := ""
		if spcSelectorStr != nil {
			s = *spcSelectorStr
		}

		spaceSelector, err = models.ParseSelector(s)
		if err != nil {
			log.Fatalf("cloud: copy command: %v", err)
		}

		console.PrintInfo(fmt.Sprintf("Using local workspace '%s'\n", spaceSelector.String()))
	}

	console.PrintCommandList("Commands to copy", commands, "static", ListingsSortOption)

	if SkipQuestionsFlag || tty.Confirm("Continue?") {
		if cloneRemoteSpace {
			ctrl.cloneSpace(cloudSelector, commands)
		} else {
			ctrl.copyCommands(spaceSelector, commands)
		}
	} else {
		console.PrintError("Cloning cancelled")
	}
}

func (ctrl *CLIController) CloudCommandView(selectorStr string) {
	selector, err := models.ParseSelectorForCloud(selectorStr)
	if err != nil {
		log.Fatalf("cloud: view command: invalid cloud selector: %v", err)
	}

	if selector.Item == "" {
		log.Fatalf("cloud: view command: command's label not specified")
	}

	command, err := ctrl.cloud.CommandFind(selector)
	if err != nil {
		log.Fatalf("cloud: view command: %v", err)
	}

	console.PrintCommand(selector.String(), command, false)
}
