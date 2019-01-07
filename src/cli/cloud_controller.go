package cli

import (
	"fmt"
	"log"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/tools"
	"github.com/dplabs/cbox/src/tools/console"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) CloudLogin(cmd *cobra.Command, args []string) {
	fmt.Println(tools.Logo)
	url := fmt.Sprintf("%s/auth/", cloud.URL)
	fmt.Printf("Open this URL in a browser and follow the authentication process: \n\n%s\n\n", url)

	jwt := console.ReadString("JWT Token", console.NOT_EMPTY_VALUES)
	fmt.Println()

	_, _, name, err := core.CloudLogin(jwt)
	if err != nil {
		console.PrintError("Error trying to parse JWT token. Try to login again")
		log.Fatalf("cloud: login: %v", err)
	}

	console.PrintSuccess("Hi " + name + "!")
}

func (ctrl *CLIController) CloudLogout(cmd *cobra.Command, args []string) {
	fmt.Println(tools.Logo)
	core.CloudLogout()
	console.PrintSuccess("Successfully logged out from cbox cloud. See you back soon!")
}
