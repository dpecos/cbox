package cli

import (
	"fmt"
	"log"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.MaximumNArgs(1),
	Short: "Add a new command to your cbox",
	Run: func(cmd *cobra.Command, args []string) {

		var selectorStr = ""
		if len(args) == 1 {
			selectorStr = args[0]
		}

		selector, err := models.ParseSelector(selectorStr)
		if err != nil {
			log.Fatal("Could not parse selector", err)
		}

		cbox := core.LoadCbox()
		space := cbox.SpaceFind(selector.Space)

		command := tools.ConsoleReadCommand()

		space.CommandAdd(command)
		core.PersistCbox(cbox)

		fmt.Println(aurora.Green("\nCommand stored successfully!\n"))
		tools.PrintCommand(command, true, false)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
