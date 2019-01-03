package cli

import (
	"fmt"
	"log"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools"
	"github.com/dplabs/cbox/src/tools/console"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) CloudCommandList(cmd *cobra.Command, args []string) {
	selector, err := models.ParseSelectorForCloudCommand(args[0])
	if err != nil {
		log.Fatalf("cloud: list commands: invalid cloud selector: %v", err)
	}

	cloud, err := core.CloudClient()
	if err != nil {
		log.Fatalf("cloud: list commands: cloud client: %v", err)
	}

	commands, err := cloud.CommandList(selector)
	if err != nil {
		log.Fatalf("cloud: list commands: %v", err)
	}

	tools.PrintCommandList(selector.String(), commands, viewSnippet, false)
}

func (ctrl *CLIController) CloudCommandCopy(cmd *cobra.Command, args []string) {
	console.PrintAction("Copying cloud commands")

	cmdSelector, err := models.ParseSelectorForCloudCommand(args[0])
	if err != nil {
		log.Fatalf("cloud: copy command: invalid cloud selector: %v", err)
	}

	spaceSelector, err := models.ParseSelectorMandatorySpace(args[1])
	if err != nil {
		log.Fatalf("cloud: copy command: invalid space selector: %v", err)
	}

	space, err := cboxInstance.SpaceFind(spaceSelector.Space)
	if err != nil {
		log.Fatalf("cloud: copy command: local space: %v", err)
	}

	cloud, err := core.CloudClient()
	if err != nil {
		log.Fatalf("cloud: copy command: cloud client: %v", err)
	}

	commands, err := cloud.CommandList(cmdSelector)
	if err != nil {
		log.Fatalf("cloud: copy command: retrieving matches: %v", err)
	}

	if len(commands) == 0 {
		console.PrintError(fmt.Sprintf("Command '%s' not found", cmdSelector))
	}

	tools.PrintCommandList("Commands to copy", commands, false, false)

	if console.Confirm(fmt.Sprintf("Copy these commands into %s?", spaceSelector)) {

		failures := false
		for _, command := range commands {
			err = space.CommandAdd(command)
			if err != nil {
				failures = true
				log.Printf("cloud: copy command: %v", err)
			}
		}

		core.Save(cboxInstance)

		if failures {
			console.PrintError("Some commands could not be stored")
		} else {
			console.PrintSuccess("Commands copied successfully!")
		}
	} else {
		console.PrintError("Copy cancelled")
	}
}
