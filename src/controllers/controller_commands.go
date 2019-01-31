package controllers

import (
	"fmt"
	"log"
	"strings"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/tools/console"
)

func (ctrl *CLIController) CommandList(args []string) {
	selector := ctrl.parseSelectorAllowEmpty(args)

	space, err := ctrl.findSpace(selector)
	if err != nil {
		log.Fatalf("list commands: %v", err)
	}

	commands := space.CommandList(selector.Item)

	console.PrintCommandList(selector.String(), commands, ShowCommandsSourceFlag, false)
}

func (ctrl *CLIController) CommandAdd(args []string) {
	console.PrintAction("Adding a new commands")

	selector := ctrl.parseSelectorAllowEmpty(args)

	space, err := ctrl.findSpace(selector)
	if err != nil {
		log.Fatalf("add command: %v", err)
	}

	fmt.Printf("Creating new command...\n")

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

func (ctrl *CLIController) CommandEdit(args []string) {
	console.PrintAction("Editing a command")

	selector := ctrl.parseSelector(args)

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

func (ctrl *CLIController) CommandDelete(args []string) {
	console.PrintAction("Deleting a command")

	selector := ctrl.parseSelector(args)

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

func (ctrl *CLIController) CommandView(args []string) {
	selector := ctrl.parseSelector(args)

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
