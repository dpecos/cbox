package controllers

import (
	"fmt"
	"log"
	"sort"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools/console"
	"github.com/dplabs/cbox/src/tools/tty"
)

func (ctrl *CLIController) TagsList(spcSelectorStr *string) {

	s := ""
	if spcSelectorStr != nil {
		s = *spcSelectorStr
	}

	selector, err := models.ParseSelector(s)
	if err != nil {
		log.Fatalf("list tags: %v", err)
	}

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

func (ctrl *CLIController) TagsAdd(cmdSelectorStr string, tags ...string) {
	console.PrintAction("Adding new tags to a command")

	selector, err := models.ParseSelector(cmdSelectorStr)
	if err != nil {
		log.Fatalf("add tags: %v", err)
	}

	space, err := ctrl.findSpace(selector)
	if err != nil {
		log.Fatalf("add tags: %v", err)
	}

	command, err := space.CommandFind(selector.Item)
	if err != nil {
		log.Fatalf("add tags: %v", err)
	}

	tty.Print("Adding tags to command with label '%s'...\n", command.Label)

	for _, tag := range tags {
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

func (ctrl *CLIController) TagsRemove(cmdSelectorStr string, tags ...string) {
	console.PrintAction("Removing tags from a command")

	selector, err := models.ParseSelector(cmdSelectorStr)
	if err != nil {
		log.Fatalf("remove tags: %v", err)
	}

	space, err := ctrl.findSpace(selector)
	if err != nil {
		log.Fatalf("remove tags: %v", err)
	}

	command, err := space.CommandFind(selector.Item)
	if err != nil {
		log.Fatalf("remove tags: %v", err)
	}

	tty.Print("Removing tags from command with label '%s'...\n", command.Label)

	for _, tag := range tags {
		if tag != "" {
			command.TagDelete(tag)
		}
	}

	core.Save(ctrl.cbox)

	console.PrintCommand("Untagged command", command, true, false)

	console.PrintSuccess("Command tag deleted successfully!")
}

func (ctrl *CLIController) TagsDelete(tagSelectorStr string) {
	console.PrintAction("Deleting tags from an space")

	selector, err := models.ParseSelectorMandatoryItem(tagSelectorStr)
	if err != nil {
		log.Fatalf("delete tag: %v", err)
	}

	space, err := ctrl.findSpace(selector)
	if err != nil {
		log.Fatalf("delete tag: %v", err)
	}
	commands := space.CommandList(selector.Item)

	tty.Print("Deleting tag from space '%s'...\n", space.String())

	for _, cmd := range commands {
		command, err := space.CommandFind(cmd.Label)
		if err != nil {
			log.Fatalf("delete tags: %v", err)
		}
		command.TagDelete(selector.Item)

		console.PrintCommand("Untagged command", command, false, false)
	}

	core.Save(ctrl.cbox)

	console.PrintSuccess(fmt.Sprintf("\nTag '%s' successfully deleted from space '%s'!", selector.Item, selector.Space))
}
