package cli

import (
	"fmt"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/tools"
	"github.com/dpecos/cbox/tools/console"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) CommandList(cmd *cobra.Command, args []string) {

	selector := parseSelectorAllowEmpty(args)

	cbox := core.LoadCbox()
	space := cbox.SpaceFind(selector.Space)
	commands := space.CommandList(selector.Item)

	for _, command := range commands {
		tools.PrintCommand(&command, viewSnippet, false)
	}
}

func (ctrl *CLIController) CommandAdd(cmd *cobra.Command, args []string) {

	selector := parseSelectorAllowEmpty(args)

	cbox := core.LoadCbox()
	space := cbox.SpaceFind(selector.Space)

	command := tools.ConsoleReadCommand()

	space.CommandAdd(command)
	core.PersistCbox(cbox)

	tools.PrintCommand(command, true, false)
	fmt.Println(aurora.Green("\nCommand stored successfully!\n"))
}

func (ctrl *CLIController) CommandEdit(cmd *cobra.Command, args []string) {

	selector := parseSelector(args)

	cbox := core.LoadCbox()

	space := cbox.SpaceFind(selector.Space)
	command := space.CommandFind(selector.Item)
	tools.ConsoleEditCommand(command)

	space.CommandEdit(command, selector.Item)

	tools.PrintCommand(command, true, false)
	if console.Confirm("Update?") {
		core.PersistCbox(cbox)
		fmt.Println(aurora.Green("\nCommand updated successfully!\n"))
	} else {
		fmt.Println("Cancelled")
	}
}

func (ctrl *CLIController) CommandDelete(cmd *cobra.Command, args []string) {

	selector := parseSelector(args)

	cbox := core.LoadCbox()
	space := cbox.SpaceFind(selector.Space)
	command := space.CommandFind(selector.Item)

	tools.PrintCommand(command, true, false)
	if console.Confirm(aurora.Red("Are you sure you want to delete this command?").String()) {
		space.CommandDelete(command)
		core.PersistCbox(cbox)
		fmt.Println(aurora.Green("\nCommand deleted successfully!\n"))
	}
}

func (ctrl *CLIController) CommandView(cmd *cobra.Command, args []string) {

	selector := parseSelector(args)

	cbox := core.LoadCbox()
	space := cbox.SpaceFind(selector.Space)
	command := space.CommandFind(selector.Item)

	tools.PrintCommand(command, true, sourceOnly)
}
