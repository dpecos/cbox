package cli

import (
	"fmt"
	"log"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/tools"
	"github.com/dpecos/cbox/tools/console"

	"github.com/spf13/cobra"
)

const (
	AUTH_URL_DEV = "https://api.dev.cbox.dplabs.io/auth/"
)

func (ctrl *CLIController) CloudLogin(cmd *cobra.Command, args []string) {
	fmt.Println(tools.Logo)
	fmt.Printf("Open this URL in a browser and follow the authentication process: \n\n%s\n\n", AUTH_URL_DEV)

	jwt := console.ReadString("JWT Token")
	fmt.Println()

	user, err := core.CloudLogin(jwt)
	if err != nil {
		console.PrintError("Error trying to parse JWT token. Try to login again")
		log.Fatalf("cloud: %v", err)
	}

	console.PrintSuccess("Hi " + user + "!")
}
