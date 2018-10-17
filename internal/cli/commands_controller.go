package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/dpecos/cbox/internal/core"
	"github.com/dpecos/cbox/tools"
	"github.com/dpecos/cbox/tools/console"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) CommandList(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelectorAllowEmpty(args)

	cbox := core.LoadCbox("")
	space, err := cbox.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("list commands: %v", err)
	}

	commands := space.CommandList(selector.Item)

	tools.PrintCommandList("", commands, viewSnippet, false)
}

func (ctrl *CLIController) CommandAdd(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelectorAllowEmpty(args)

	cbox := core.LoadCbox("")
	space, err := cbox.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("add command: %v", err)
	}

	fmt.Println("Creating new command")
	command := tools.ConsoleReadCommand()

	err = space.CommandAdd(command)
	for err != nil {
		console.PrintError(fmt.Sprintf("\nLabel '%s' already found in space. Try a different one", command.Label))
		command.Label = strings.ToLower(console.ReadString("Label"))
		err = space.CommandAdd(command)
	}
	core.PersistCbox(cbox)

	tools.PrintCommand("New command", command, true, false)

	console.PrintSuccess("Command stored successfully!")
}

func (ctrl *CLIController) CommandEdit(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelector(args)

	cbox := core.LoadCbox("")

	space, err := cbox.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("edit command: %v", err)
	}

	command, err := space.CommandFind(selector.Item)
	if err != nil {
		log.Fatalf("edit command: %v", err)
	}

	previousCommandLabel := command.Label

	tools.PrintCommand("Command to edit", command, true, false)

	tools.ConsoleEditCommand(command)

	err = space.CommandEdit(command, previousCommandLabel)
	for err != nil {
		console.PrintError(fmt.Sprintf("Label '%s' already found in space. Try a different one", command.Label))
		command.Label = strings.ToLower(console.ReadString("Label"))
		err = space.CommandEdit(command, previousCommandLabel)
	}

	tools.PrintCommand("Command after edition", command, true, false)

	if console.Confirm("Update?") {
		core.PersistCbox(cbox)
		console.PrintSuccess("Command updated successfully!")
	} else {
		console.PrintError("Edition cancelled")
	}
}

func (ctrl *CLIController) CommandDelete(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelector(args)

	cbox := core.LoadCbox("")
	space, err := cbox.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("delete command: %v", err)
	}

	command, err := space.CommandFind(selector.Item)
	if err != nil {
		log.Fatalf("delete command: %v", err)
	}

	tools.PrintCommand("Command to delete ", command, true, false)

	if console.Confirm("Are you sure you want to delete this command?") {
		space.CommandDelete(command)
		core.PersistCbox(cbox)
		console.PrintSuccess("Command deleted successfully!")
	} else {
		console.PrintError("Deletion cancelled")
	}
}

func (ctrl *CLIController) CommandView(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelector(args)

	cbox := core.LoadCbox("")
	space, err := cbox.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("view command: %v", err)
	}

	command, err := space.CommandFind(selector.Item)
	if err != nil {
		log.Fatalf("view command: %v", err)
	}

	tools.PrintCommand("", command, true, sourceOnly)
}
