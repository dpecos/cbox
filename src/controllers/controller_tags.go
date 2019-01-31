package controllers

import (
	"fmt"
	"log"
	"sort"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools/console"
)

func (ctrl *CLIController) TagsList(args []string) {
	selector := ctrl.parseSelectorAllowEmpty(args)

	space, err := ctrl.findSpace(selector)
	if err != nil {
		log.Fatalf("list tags: %v", err)
	}

	tags := space.TagsList(selector.Item)
	sort.Strings(tags)

	for _, tag := range tags {
		console.PrintTag(tag)
	}
}

func (ctrl *CLIController) TagsAdd(args []string) {
	console.PrintAction("Adding new tags to a command")

	selector := ctrl.parseSelector(args)

	space, err := ctrl.findSpace(selector)
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

	core.Save(ctrl.cbox)

	console.PrintCommand("Tagged command", command, true, false)

	console.PrintSuccess("Command tagged successfully!")
}

func (ctrl *CLIController) TagsRemove(args []string) {
	console.PrintAction("Removing tags from a command")

	selector := ctrl.parseSelector(args)

	space, err := ctrl.findSpace(selector)
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

	core.Save(ctrl.cbox)

	console.PrintCommand("Untagged command", command, true, false)

	console.PrintSuccess("Command tag deleted successfully!")
}

func (ctrl *CLIController) TagsDelete(args []string) {
	console.PrintAction("Deleting tags from an space")

	selector, err := models.ParseSelectorMandatoryItem(args[0])
	if err != nil {
		log.Fatalf("delete tags: %v", err)
	}

	space, err := ctrl.findSpace(selector)
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

		console.PrintCommand("Untagged command", command, false, false)
	}

	core.Save(ctrl.cbox)

	msg := fmt.Sprintf("\nTag '%s' successfully deleted from space '%s'!", selector.Item, selector.Space)
	console.PrintSuccess(msg)
}
