package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/dpecos/cbox/internal/app/core"
	"github.com/dpecos/cbox/internal/pkg"
	"github.com/dpecos/cbox/internal/pkg/console"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) CommandList(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelectorAllowEmpty(args)

	space, err := cboxInstance.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("list commands: %v", err)
	}

	commands := space.CommandList(selector.Item)

	pkg.PrintCommandList("", commands, viewSnippet, false)
}

func (ctrl *CLIController) CommandAdd(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelectorAllowEmpty(args)

	space, err := cboxInstance.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("add command: %v", err)
	}

	fmt.Println("Creating new command")
	command := pkg.ConsoleReadCommand()

	err = space.CommandAdd(command)
	for err != nil {
		console.PrintError(fmt.Sprintf("\nLabel '%s' already found in space. Try a different one", command.Label))
		command.Label = strings.ToLower(console.ReadString("Label"))
		err = space.CommandAdd(command)
	}
	core.Save(cboxInstance)

	pkg.PrintCommand("New command", command, true, false)

	console.PrintSuccess("Command stored successfully!")
}

func (ctrl *CLIController) CommandEdit(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelector(args)

	space, err := cboxInstance.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("edit command: %v", err)
	}

	command, err := space.CommandFind(selector.Item)
	if err != nil {
		log.Fatalf("edit command: %v", err)
	}

	previousCommandLabel := command.Label

	pkg.PrintCommand("Command to edit", command, true, false)

	pkg.ConsoleEditCommand(command)

	err = space.CommandEdit(command, previousCommandLabel)
	for err != nil {
		console.PrintError(fmt.Sprintf("Label '%s' already found in space. Try a different one", command.Label))
		command.Label = strings.ToLower(console.ReadString("Label"))
		err = space.CommandEdit(command, previousCommandLabel)
	}

	pkg.PrintCommand("Command after edition", command, true, false)

	if console.Confirm("Update?") {
		core.Save(cboxInstance)
		console.PrintSuccess("Command updated successfully!")
	} else {
		console.PrintError("Edition cancelled")
	}
}

func (ctrl *CLIController) CommandDelete(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelector(args)

	space, err := cboxInstance.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("delete command: %v", err)
	}

	command, err := space.CommandFind(selector.Item)
	if err != nil {
		log.Fatalf("delete command: %v", err)
	}

	pkg.PrintCommand("Command to delete ", command, true, false)

	if console.Confirm("Are you sure you want to delete this command?") {
		space.CommandDelete(command)
		core.Save(cboxInstance)
		console.PrintSuccess("Command deleted successfully!")
	} else {
		console.PrintError("Deletion cancelled")
	}
}

func (ctrl *CLIController) CommandView(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelector(args)

	space, err := cboxInstance.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("view command: %v", err)
	}

	command, err := space.CommandFind(selector.Item)
	if err != nil {
		log.Fatalf("view command: %v", err)
	}

	pkg.PrintCommand("", command, true, sourceOnly)
}
