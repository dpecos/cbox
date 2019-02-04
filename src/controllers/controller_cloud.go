package controllers

import (
	"fmt"
	"log"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/tools"
	"github.com/dplabs/cbox/src/tools/console"
)

func (ctrl *CLIController) CloudLogin() {
	fmt.Println(tools.Logo)
	url := fmt.Sprintf("%s/auth/", ctrl.cloud.URL)
	fmt.Printf("Open this URL in a browser and follow the authentication process: \n\n%s\n\n", url)

	jwt := console.ReadString("JWT Token", console.NOT_EMPTY_VALUES)
	fmt.Println()

	_, _, name, err := core.CloudLogin(jwt)
	if err != nil {
		console.PrintError("Error trying to parse JWT token. Try to login again")
		log.Fatalf("ctrl.cloud: login: %v", err)
	}

	console.PrintSuccess("Hi " + name + "!")
}

func (ctrl *CLIController) CloudLogout() {
	fmt.Println(tools.Logo)
	core.CloudLogout()
	console.PrintSuccess("Successfully logged out from cbox ctrl.cloud. See you back soon!")
}
