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

var tagDeleteCmd = &cobra.Command{
	Use:   "delete",
	Args:  cobra.MinimumNArgs(2),
	Short: "Delete tags from a command",
	Run: func(cmd *cobra.Command, args []string) {

		selector, err := models.ParseSelector(args[0])
		if err != nil {
			log.Fatal("Could not parse selector", err)
		}

		cbox := core.LoadCbox()

		space := cbox.SpaceFind(selector.Space)
		command := space.CommandFind(selector.Item)

		for _, tag := range args[1:] {
			command.TagDelete(tag)
		}

		core.PersistCbox(cbox)

		tools.PrintCommand(command, true, false)
		fmt.Println(aurora.Green("\nCommand tag deleted successfully!\n"))
	},
}

func init() {
	tagCmd.AddCommand(tagDeleteCmd)
}
