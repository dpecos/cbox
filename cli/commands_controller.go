package cli

import (
	"fmt"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/tools"
	"github.com/dpecos/cbox/tools/console"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) CommandList(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelectorAllowEmpty(args)

	cbox := core.LoadCbox()
	space := cbox.SpaceFind(selector.Space)
	commands := space.CommandList(selector.Item)

	for _, command := range commands {
		tools.PrintCommand(&command, viewSnippet, false)
	}
}

func (ctrl *CLIController) CommandAdd(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelectorAllowEmpty(args)

	cbox := core.LoadCbox()
	space := cbox.SpaceFind(selector.Space)

	fmt.Println("Creating new command")
	command := tools.ConsoleReadCommand()

	space.CommandAdd(command)
	core.PersistCbox(cbox)

	fmt.Println("\n--- New command ---")
	tools.PrintCommand(command, true, false)
	fmt.Println("-----\n")

	console.PrintSuccess("Command stored successfully!")
}

func (ctrl *CLIController) CommandEdit(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelector(args)

	cbox := core.LoadCbox()

	space := cbox.SpaceFind(selector.Space)
	command := space.CommandFind(selector.Item)

	fmt.Printf("Editing command with ID %s\n", command.ID)
	tools.ConsoleEditCommand(command)

	space.CommandEdit(command, selector.Item)

	fmt.Println("\n--- Command after edited values ---")
	tools.PrintCommand(command, true, false)
	fmt.Println("-----\n")

	if console.Confirm("Update?") {
		core.PersistCbox(cbox)
		console.PrintSuccess("\nCommand updated successfully!")
	} else {
		console.PrintError("Edition cancelled")
	}
}

func (ctrl *CLIController) CommandDelete(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelector(args)

	cbox := core.LoadCbox()
	space := cbox.SpaceFind(selector.Space)
	command := space.CommandFind(selector.Item)

	fmt.Println("\n--- Command to delete ---")
	tools.PrintCommand(command, true, false)
	fmt.Println("-----\n")

	if console.Confirm("Are you sure you want to delete this command?") {
		space.CommandDelete(command)
		core.PersistCbox(cbox)
		console.PrintSuccess("\nCommand deleted successfully!")
	} else {
		console.PrintError("Deletion cancelled")
	}
}

func (ctrl *CLIController) CommandView(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelector(args)

	cbox := core.LoadCbox()
	space := cbox.SpaceFind(selector.Space)
	command := space.CommandFind(selector.Item)

	tools.PrintCommand(command, true, sourceOnly)
}
