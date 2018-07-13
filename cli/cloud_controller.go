package cli

import (
	"fmt"
	"log"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools"
	"github.com/dpecos/cbox/tools/console"

	"github.com/spf13/cobra"
)

func (ctrl *CLIController) CloudLogin(cmd *cobra.Command, args []string) {
	fmt.Println(tools.Logo)
	fmt.Printf("Open this URL in a browser and follow the authentication process: \n\n%s\n\n", fmt.Sprintf("%s/auth/", core.SERVER_URL_DEV))

	jwt := console.ReadString("JWT Token")
	fmt.Println()

	user, err := core.CloudLogin(jwt)
	if err != nil {
		console.PrintError("Error trying to parse JWT token. Try to login again")
		log.Fatalf("cloud: %v", err)
	}

	console.PrintSuccess("Hi " + user + "!")
}

func (ctrl *CLIController) CloudLogout(cmd *cobra.Command, args []string) {
	fmt.Println(tools.Logo)
	core.CloudLogout()
	console.PrintSuccess("Successfully logged out from cbox cloud. See you back soon!")
}

func (ctrl *CLIController) CloudPublishSpace(cmd *cobra.Command, args []string) {
	selector, err := models.ParseSelectorMandatorySpace(args[0])
	if err != nil {
		log.Fatalf("publish space: %v", err)
	}

	cbox := core.LoadCbox()

	space := cbox.SpaceFind(selector.Space)

	fmt.Printf("--- Space ---\n")
	tools.PrintSpace(space)
	fmt.Printf("-----\n\n")

	if console.Confirm("Publish?") {
		fmt.Printf("Publishing space '%s'...\n", space.ID)

		cloud, err := core.CloudClient()
		if err != nil {
			log.Fatalf("cloud: %v", err)
		}
		err = cloud.PublishSpace(space)
		if err != nil {
			log.Fatalf("cloud: %v", err)
		}

		console.PrintSuccess("Space published successfully!")
	} else {
		console.PrintError("Publish cancelled")
	}
}
