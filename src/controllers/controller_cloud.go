package controllers

import (
	"fmt"
	"log"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/tools"
	"github.com/dplabs/cbox/src/tools/console"
	"github.com/dplabs/cbox/src/tools/tty"
)

func (ctrl *CLIController) CloudLogin() {
	tty.Print(tools.Logo + "\n")
	url := fmt.Sprintf("%s/auth/", ctrl.cloud.URL)
	tty.Print("Open this URL in a browser and follow the authentication process: \n\n%s\n\n", url)

	jwt := console.ReadString("JWT Token", console.NOT_EMPTY_VALUES)
	tty.Print("\n")

	name, err := ctrl.cloud.ServerLogin(jwt)
	if err != nil {
		console.PrintError("Error trying to parse JWT token. Try to login again")
		log.Fatalf("cloud: login: %v", err)
	}
	core.StoreCloudSettings(ctrl.cloud)

	console.PrintSuccess("Hi " + name + "!")
}

func (ctrl *CLIController) CloudLogout() {
	tty.Print(tools.Logo + "\n")
	core.DeleteCloudSettings()
	console.PrintSuccess("Successfully logged out from cbox cloud. See you back soon!")
}
