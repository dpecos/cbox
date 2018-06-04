package cli

import (
	"fmt"
	"log"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools"
	"github.com/dpecos/cbox/tools/console"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Args:  cobra.ExactArgs(1),
	Short: "Edit a command from your cbox",
	Run: func(cmd *cobra.Command, args []string) {

		selector, err := models.ParseSelector(args[0])
		if err != nil {
			log.Fatal("Could not parse selector", err)
		}

		cbox := core.LoadCbox()

		space := cbox.SpaceFind(selector.Space)
		command := space.CommandFind(selector.Item)
		tools.ConsoleEditCommand(command)

		space.CommandEdit(command, selector.Item)

		tools.PrintCommand(command, true, false)
		if console.Confirm("Update?") {
			core.PersistCbox(cbox)
			fmt.Println(aurora.Green("\nCommand updated successfully!\n"))
		} else {
			fmt.Println("Cancelled")
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
