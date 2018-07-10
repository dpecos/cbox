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
		if viewSnippet {
			fmt.Printf("\n------------------\n\n")
		}
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

	fmt.Printf("\n--- New command ---\n")
	tools.PrintCommand(command, true, false)
	fmt.Printf("-----\n\n")

	console.PrintSuccess("Command stored successfully!")
}

func (ctrl *CLIController) CommandEdit(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelector(args)

	cbox := core.LoadCbox()

	space := cbox.SpaceFind(selector.Space)
	command := space.CommandFind(selector.Item)

	fmt.Printf("\n--- Command to edit ---\n")
	tools.PrintCommand(command, true, false)
	fmt.Printf("-----\n\n")

	tools.ConsoleEditCommand(command)

	space.CommandEdit(command, selector.Item)

	fmt.Printf("\n--- Command after edited values ---\n")
	tools.PrintCommand(command, true, false)
	fmt.Printf("-----\n\n")

	if console.Confirm("Update?") {
		core.PersistCbox(cbox)
		console.PrintSuccess("Command updated successfully!")
	} else {
		console.PrintError("Edition cancelled")
	}
}

func (ctrl *CLIController) CommandDelete(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelector(args)

	cbox := core.LoadCbox()
	space := cbox.SpaceFind(selector.Space)
	command := space.CommandFind(selector.Item)

	fmt.Printf("\n--- Command to delete ---\n")
	tools.PrintCommand(command, true, false)
	fmt.Printf("-----\n\n")

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

	cbox := core.LoadCbox()
	space := cbox.SpaceFind(selector.Space)
	command := space.CommandFind(selector.Item)

	tools.PrintCommand(command, true, sourceOnly)
}
