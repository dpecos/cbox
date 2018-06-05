package cli

import (
	"fmt"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/tools"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var spacesAddCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.ExactArgs(0),
	Short: "Add a new space to your cbox",
	Run: func(cmd *cobra.Command, args []string) {

		cbox := core.LoadCbox()

		space := tools.ConsoleReadSpace()
		cbox.SpaceCreate(space)

		core.PersistCbox(cbox)

		fmt.Println(aurora.Green("\nSpace successfully created!\n"))
		tools.PrintSpace(space)
	},
}

func init() {
	spacesCmd.AddCommand(spacesAddCmd)
}
