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

func (ctrl *CLIController) CommandList(spcSelectorStr *string) {
	s := ""
	if spcSelectorStr != nil {
		s = *spcSelectorStr
	}

	selector, err := models.ParseSelector(s)
	if err != nil {
		log.Fatalf("list commands: %v", err)
	}

	space, err := ctrl.findSpace(selector)
	if err != nil {
		log.Fatalf("list commands: %v", err)
	}

	commands := space.CommandList(selector.Item)

	console.PrintCommandList(selector.String(), commands, ShowCommandsSourceFlag, false)
}

func (ctrl *CLIController) CommandAdd(spcSelectorStr *string) {
	console.PrintAction("Adding a new command")

	s := ""
	if spcSelectorStr != nil {
		s = *spcSelectorStr
	}

	selector, err := models.ParseSelector(s)
	if err != nil {
		log.Fatalf("add command: %v", err)
	}

	space, err := ctrl.findSpace(selector)
	if err != nil {
		log.Fatalf("add command: %v", err)
	}

	tty.Print("Data for new command:\n")

	command := console.ReadCommand(space)

	err = space.CommandAdd(command)
	for err != nil {
		console.PrintError(fmt.Sprintf("Label '%s' already found in space. Try a different one", command.Label))
		command.Label = strings.ToLower(console.ReadString("Label", console.NOT_EMPTY_VALUES, console.ONLY_VALID_CHARS))
		err = space.CommandAdd(command)
	}
	core.Save(ctrl.cbox)

	console.PrintCommand("New command", command, true, false)

	console.PrintSuccess("Command stored successfully!")
}

func (ctrl *CLIController) CommandEdit(spcSelectorStr string) {
	console.PrintAction("Editing a command")

	selector, err := models.ParseSelector(spcSelectorStr)
	if err != nil {
		log.Fatalf("edit command: %v", err)
	}

	space, err := ctrl.findSpace(selector)
	if err != nil {
		log.Fatalf("edit command: %v", err)
	}

	command, err := space.CommandFind(selector.Item)
	if err != nil {
		log.Fatalf("edit command: %v", err)
	}

	previousCommandLabel := command.Label

	console.PrintCommand("Command to edit", command, true, false)

	console.EditCommand(command)
	command.Selector.Item = command.Label

	err = space.CommandEdit(command, previousCommandLabel)
	for err != nil {
		console.PrintError(fmt.Sprintf("Label '%s' already found in space. Try a different one", command.Label))
		command.Label = strings.ToLower(console.ReadString("Label", console.NOT_EMPTY_VALUES, console.ONLY_VALID_CHARS))
		command.Selector.Item = command.Label
		err = space.CommandEdit(command, previousCommandLabel)
	}

	console.PrintCommand("Command after edition", command, true, false)

	if SkipQuestionsFlag || console.Confirm("Update?") {
		core.Save(ctrl.cbox)
		console.PrintSuccess("Command updated successfully!")
	} else {
		console.PrintError("Edition cancelled")
	}
}

func (ctrl *CLIController) CommandDelete(spcSelectorStr string) {
	console.PrintAction("Deleting a command")

	selector, err := models.ParseSelector(spcSelectorStr)
	if err != nil {
		log.Fatalf("delete command: %v", err)
	}

	space, err := ctrl.findSpace(selector)
	if err != nil {
		log.Fatalf("delete command: %v", err)
	}

	command, err := space.CommandFind(selector.Item)
	if err != nil {
		log.Fatalf("delete command: %v", err)
	}

	console.PrintCommand("Command to delete ", command, true, false)

	if SkipQuestionsFlag || console.Confirm("Are you sure you want to delete this command?") {
		space.CommandDelete(command)
		core.Save(ctrl.cbox)
		console.PrintSuccess("Command deleted successfully!")
	} else {
		console.PrintError("Deletion cancelled")
	}
}

func (ctrl *CLIController) CommandView(spcSelectorStr string) {
	selector, err := models.ParseSelector(spcSelectorStr)
	if err != nil {
		log.Fatalf("view command: %v", err)
	}

	space, err := ctrl.findSpace(selector)
	if err != nil {
		log.Fatalf("view command: %v", err)
	}

	command, err := space.CommandFind(selector.Item)
	if err != nil {
		log.Fatalf("view command: %v", err)
	}

	console.PrintCommand("", command, true, SourceOnlyFlag)
}
