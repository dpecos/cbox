package cli

import (
	"fmt"
	"log"
	"sort"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools"
	"github.com/dplabs/cbox/src/tools/console"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) TagsList(cmd *cobra.Command, args []string) {
	selector := ctrl.parseSelectorAllowEmpty(args)

	space, err := findSpace(selector)
	if err != nil {
		log.Fatalf("list tags: %v", err)
	}

	tags := space.TagsList(selector.Item)
	sort.Strings(tags)

	for _, tag := range tags {
		tools.PrintTag(tag)
	}
}

func (ctrl *CLIController) TagsAdd(cmd *cobra.Command, args []string) {
	console.PrintAction("Adding new tags to a command")

	selector := ctrl.parseSelector(args)

	space, err := findSpace(selector)
	if err != nil {
		log.Fatalf("add tags: %v", err)
	}

	command, err := space.CommandFind(selector.Item)
	if err != nil {
		log.Fatalf("add tags: %v", err)
	}

	fmt.Printf("Adding tags to command with label '%s'...\n", command.Label)

	for _, tag := range args[1:] {
		if tag != "" {
			if !console.CheckValidChars(tag) {
				log.Fatalf("add tags: invalid characters in tag '%s'", tag)
			}
			command.TagAdd(tag)
		}
	}

	core.Save(cboxInstance)

	tools.PrintCommand("Tagged command", command, true, false)

	console.PrintSuccess("Command tagged successfully!")
}

func (ctrl *CLIController) TagsRemove(cmd *cobra.Command, args []string) {
	console.PrintAction("Removing tags from a command")

	selector := ctrl.parseSelector(args)

	space, err := findSpace(selector)
	if err != nil {
		log.Fatalf("remove tags: %v", err)
	}

	command, err := space.CommandFind(selector.Item)
	if err != nil {
		log.Fatalf("remove tags: %v", err)
	}

	fmt.Printf("Removing tags from command with label '%s'...\n", command.Label)

	for _, tag := range args[1:] {
		if tag != "" {
			command.TagDelete(tag)
		}
	}

	core.Save(cboxInstance)

	tools.PrintCommand("Untagged command", command, true, false)

	console.PrintSuccess("Command tag deleted successfully!")
}

func (ctrl *CLIController) TagsDelete(cmd *cobra.Command, args []string) {
	console.PrintAction("Deleting tags from an space")

	selector, err := models.ParseSelectorMandatoryItem(args[0])
	if err != nil {
		log.Fatalf("delete tags: %v", err)
	}

	space, err := findSpace(selector)
	if err != nil {
		log.Fatalf("delete tags: %v", err)
	}
	commands := space.CommandList(selector.Item)

	fmt.Printf("Deleting tags from space '%s'...\n", space.String())

	for _, cmd := range commands {
		command, err := space.CommandFind(cmd.Label)
		if err != nil {
			log.Fatalf("delete tags: %v", err)
		}
		command.TagDelete(selector.Item)

		tools.PrintCommand("Untagged command", command, false, false)
	}

	core.Save(cboxInstance)

	msg := fmt.Sprintf("\nTag '%s' successfully deleted from space '%s'!", selector.Item, selector.Space)
	console.PrintSuccess(msg)
}
