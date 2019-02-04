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

	name, err := ctrl.cloud.ServerLogin(jwt)
	if err != nil {
		console.PrintError("Error trying to parse JWT token. Try to login again")
		log.Fatalf("ctrl.cloud: login: %v", err)
	}
	core.StoreCloudSettings(ctrl.cloud)

	console.PrintSuccess("Hi " + name + "!")
}

func (ctrl *CLIController) CloudLogout() {
	fmt.Println(tools.Logo)
	core.DeleteCloudSettings()
	console.PrintSuccess("Successfully logged out from cbox cloud. See you back soon!")
}
