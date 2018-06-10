package cli

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/dpecos/cbox/tools/console"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) TagsList(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelectorAllowEmpty(args)

	cbox := core.LoadCbox()
	space := cbox.SpaceFind(selector.Space)

	tags := space.TagsList(selector.Item)
	sort.Strings(tags)

	for _, tag := range tags {
		fmt.Printf("%s\n", tag)
	}
}

func (ctrl *CLIController) TagsAdd(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelector(args)

	cbox := core.LoadCbox()

	space := cbox.SpaceFind(selector.Space)
	command := space.CommandFind(selector.Item)

	fmt.Println("Adding tags to command with ID %s", command.ID)

	for _, tag := range args[1:] {
		command.TagAdd(strings.ToLower(tag))
	}

	core.PersistCbox(cbox)

	fmt.Println("\n--- Tagged command ---")
	tools.PrintCommand(command, true, false)
	fmt.Println("-----\n")

	console.PrintSuccess("Command tagged successfully!")
}

func (ctrl *CLIController) TagsRemove(cmd *cobra.Command, args []string) {

	selector := ctrl.parseSelector(args)

	cbox := core.LoadCbox()

	space := cbox.SpaceFind(selector.Space)
	command := space.CommandFind(selector.Item)

	fmt.Println("Removing tags from command with ID %s", command.ID)

	for _, tag := range args[1:] {
		command.TagDelete(tag)
	}

	core.PersistCbox(cbox)

	fmt.Println("\n--- Untagged command ---")
	tools.PrintCommand(command, true, false)
	fmt.Println("-----\n")

	console.PrintSuccess("Command tag deleted successfully!")
}

func (ctrl *CLIController) TagsDelete(cmd *cobra.Command, args []string) {

	selector, err := models.ParseSelectorMandatoryItem(args[0])
	if err != nil {
		log.Fatal("Could not parse selector", err)
	}

	cbox := core.LoadCbox()
	space := cbox.SpaceFind(selector.Space)
	commands := space.CommandList(selector.Item)

	fmt.Println("Deleting tags from space with Name %s", space.Name)

	for _, cmd := range commands {
		command := space.CommandFind(cmd.ID)
		command.TagDelete(selector.Item)

		fmt.Println("\n--- Untagged command ---")
		tools.PrintCommand(command, false, false)
		fmt.Println("-----\n")
	}

	core.PersistCbox(cbox)

	msg := fmt.Sprintf("\nTag '%s' successfully deleted from space '%s'!", selector.Item, selector.Space)
	console.PrintSuccess(msg)
}
