package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/dplabs/cbox/internal/app/core"
	"github.com/dplabs/cbox/internal/pkg"
	"github.com/dplabs/cbox/internal/pkg/console"
	"github.com/dplabs/cbox/pkg/models"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) CloudSpaceInfo(cmd *cobra.Command, args []string) {
	console.PrintAction("Retrieving info of an space")

	selector, err := models.ParseSelectorMandatorySpace(args[0])
	if err != nil {
		log.Fatalf("cloud: space info: %v", err)
	}

	cloud, err := core.CloudClient()
	if err != nil {
		log.Fatalf("cloud: space info: %v", err)
	}
	space, err := cloud.SpaceFind(selector, nil)
	if err != nil {
		log.Fatalf("cloud: space info: %v", err)
	}

	pkg.PrintSpace(fmt.Sprintf("Info of remote space %s", selector), space)
}

func (ctrl *CLIController) CloudSpacePublish(cmd *cobra.Command, args []string) {
	console.PrintAction("Publishing an space")

	selector, err := models.ParseSelectorMandatorySpace(args[0])
	if err != nil {
		log.Fatalf("cloud: publish space: %v", err)
	}

	space, err := cboxInstance.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("cloud: publish space: %v", err)
	}

	pkg.PrintSpace("Space to publish", space)

	if selector.Item != "" {
		commands := space.CommandList(selector.Item)
		if len(commands) == 0 {
			log.Fatalf("cloud: no local commands matched selector: %s", selector)
		}

		space.Entries = commands
	}

	// pkg.PrintCommandList("Containing these commands", space.Entries, false, false)

	if console.Confirm("Publish?") {
		fmt.Printf("Publishing space '%s'...\n", space.Label)

		cloud, err := core.CloudClient()
		if err != nil {
			log.Fatalf("cloud: publish space: %v", err)
		}
		err = cloud.SpacePublish(space)
		if err != nil {
			log.Fatalf("cloud: publish space: %v", err)
		}

		console.PrintSuccess("Space published successfully!")
	} else {
		console.PrintError("Publishing cancelled")
	}
}

func (ctrl *CLIController) CloudSpaceUnpublish(cmd *cobra.Command, args []string) {
	console.PrintAction("Unpublishing an space")

	selector, err := models.ParseSelectorMandatorySpace(args[0])
	if err != nil {
		log.Fatalf("cloud: unpublish space: %v", err)
	}

	pkg.PrintSelector("Space to unpublish", selector)

	_, err = cboxInstance.SpaceFind(selector.Space)
	if err == nil {
		console.PrintInfo("Local copy won't be deleted")
	} else {
		console.PrintWarning("You don't have a local copy of the space")
	}

	if console.Confirm("Unpublish?") {
		fmt.Printf("Unpublishing space '%s'...\n", selector.String())

		cloud, err := core.CloudClient()
		if err != nil {
			log.Fatalf("cloud: unpublish space: %v", err)
		}
		err = cloud.SpaceUnpublish(selector)
		if err != nil {
			log.Fatalf("cloud: unpublish space: %v", err)
		}

		console.PrintSuccess("Space unpublished successfully!")
	} else {
		console.PrintError("Unpublishing cancelled")
	}
}

func (ctrl *CLIController) CloudSpaceClone(cmd *cobra.Command, args []string) {
	console.PrintAction("Cloning an space")

	selector, err := models.ParseSelectorForCloudCommand(args[0])
	if err != nil {
		log.Fatalf("cloud: clone space: invalid cloud selector: %v", err)
	}

	cloud, err := core.CloudClient()
	if err != nil {
		log.Fatalf("cloud: clone space: cloud client: %v", err)
	}

	space, err := cloud.SpaceFind(selector, nil)
	if err != nil {
		log.Fatalf("cloud: clone space: %v", err)
	}

	commands, err := cloud.CommandList(selector)
	if err != nil {
		log.Fatalf("cloud: list commands: %v", err)
	}

	space.Entries = commands

	pkg.PrintSpace("Space to clone", space)
	pkg.PrintCommandList("Containing these commands", space.Entries, false, false)

	if console.Confirm("Clone?") {
		err := cboxInstance.SpaceCreate(space)
		for err != nil {
			console.PrintError("Space already found in your cbox. Try a different one")
			space.Label = strings.ToLower(console.ReadString("Label", console.NOT_EMPTY_VALUES, console.ONLY_VALID_CHARS))
			err = cboxInstance.SpaceCreate(space)
		}

		core.Save(cboxInstance)

		console.PrintSuccess("Space cloned successfully!")
	} else {
		console.PrintError("Clone cancelled")
	}
}

func (ctrl *CLIController) CloudSpacePull(cmd *cobra.Command, args []string) {
	console.PrintAction("Pulling latest changes of an space")

	selector, err := models.ParseSelectorMandatorySpace(args[0])
	if err != nil {
		log.Fatalf("cloud: pull space: invalid cloud selector: %v", err)
	}

	cloud, err := core.CloudClient()
	if err != nil {
		log.Fatalf("cloud: pull space: cloud client: %v", err)
	}

	space, err := cboxInstance.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("cloud: pull space: %v", err)
	}

	spaceCloud, err := cloud.SpaceFind(nil, &space.ID)
	if err != nil {
		log.Fatalf("cloud: pull space: %v", err)
	}

	commands, err := cloud.CommandList(selector)
	if err != nil {
		log.Fatalf("cloud: list commands: %v", err)
	}

	// Note: Label is not overwritten because user can renamed his local copy of the space
	space.Entries = commands
	space.UpdatedAt = spaceCloud.UpdatedAt
	space.Description = spaceCloud.Description

	core.Save(cboxInstance)

	pkg.PrintSpace("Pulled space", space)

	console.PrintSuccess("Space pulled successfully!")
}
