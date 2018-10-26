package cli

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/dpecos/cbox/internal/app/core"
	"github.com/dpecos/cbox/internal/pkg"
	"github.com/dpecos/cbox/internal/pkg/console"
	"github.com/dpecos/cbox/pkg/models"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) TagsList(cmd *cobra.Command, args []string) {
	selector := ctrl.parseSelectorAllowEmpty(args)

	space, err := cboxInstance.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("list tags: %v", err)
	}

	tags := space.TagsList(selector.Item)
	sort.Strings(tags)

	for _, tag := range tags {
		pkg.PrintTag(tag)
	}
}

func (ctrl *CLIController) TagsAdd(cmd *cobra.Command, args []string) {
	console.PrintAction("Adding new tags to a command")

	selector := ctrl.parseSelector(args)

	space, err := cboxInstance.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("add tags: %v", err)
	}

	command, err := space.CommandFind(selector.Item)
	if err != nil {
		log.Fatalf("add tags: %v", err)
	}

	console.PrintAction(fmt.Sprintf("Adding tags to command with label '%s'", command.Label))

	for _, tag := range args[1:] {
		if tag != "" {
			command.TagAdd(strings.ToLower(tag))
		}
	}

	core.Save(cboxInstance)

	pkg.PrintCommand("Tagged command", command, true, false)

	console.PrintSuccess("Command tagged successfully!")
}

func (ctrl *CLIController) TagsRemove(cmd *cobra.Command, args []string) {
	console.PrintAction("Removing tags from a command")

	selector := ctrl.parseSelector(args)

	space, err := cboxInstance.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("remove tags: %v", err)
	}

	command, err := space.CommandFind(selector.Item)
	if err != nil {
		log.Fatalf("remove tags: %v", err)
	}

	console.PrintAction(fmt.Sprintf("Removing tags from command with label '%s'", command.Label))

	for _, tag := range args[1:] {
		if tag != "" {
			command.TagDelete(tag)
		}
	}

	core.Save(cboxInstance)

	pkg.PrintCommand("Untagged command", command, true, false)

	console.PrintSuccess("Command tag deleted successfully!")
}

func (ctrl *CLIController) TagsDelete(cmd *cobra.Command, args []string) {
	console.PrintAction("Deleting tags from an space")

	selector, err := models.ParseSelectorMandatoryItem(args[0])
	if err != nil {
		log.Fatalf("delete tags: %v", err)
	}

	space, err := cboxInstance.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("delete tags: %v", err)
	}
	commands := space.CommandList(selector.Item)

	console.PrintAction(fmt.Sprintf("Deleting tags from space '%s'", space.Label))

	for _, cmd := range commands {
		command, err := space.CommandFind(cmd.Label)
		if err != nil {
			log.Fatalf("delete tags: %v", err)
		}
		command.TagDelete(selector.Item)

		pkg.PrintCommand("Untagged command", command, false, false)
	}

	core.Save(cboxInstance)

	msg := fmt.Sprintf("\nTag '%s' successfully deleted from space '%s'!", selector.Item, selector.Space)
	console.PrintSuccess(msg)
}
