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

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Args:    cobra.ExactArgs(1),
	Short:   "Delete a command from your cbox",
	Long:    tools.Logo,
	Run: func(cmd *cobra.Command, args []string) {

		selector, err := models.ParseSelector(args[0])
		if err != nil {
			log.Fatal("Could not parse selector", err)
		}

		cbox := core.LoadCbox()
		space := cbox.SpaceFind(selector.Space)
		command := space.CommandFind(selector.Item)

		tools.PrintCommand(command, true, false)
		if console.Confirm(aurora.Red("Are you sure you want to delete this command?").String()) {
			space.CommandDelete(command)
			core.PersistCbox(cbox)
			fmt.Println(aurora.Green("\nCommand deleted successfully!\n"))
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
