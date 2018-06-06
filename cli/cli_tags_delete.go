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

var tagsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Args:  cobra.ExactArgs(1),
	Short: "Delete specified tag",
	Long:  tools.Logo,
	Run: func(cmd *cobra.Command, args []string) {

		selector, err := models.ParseSelectorMandatoryItem(args[0])
		if err != nil {
			log.Fatal("Could not parse selector", err)
		}

		cbox := core.LoadCbox()
		space := cbox.SpaceFind(selector.Space)
		commands := space.CommandList(selector.Item)

		for _, cmd := range commands {
			command := space.CommandFind(cmd.ID)
			command.TagDelete(selector.Item)
			tools.PrintCommand(command, false, false)
		}

		core.PersistCbox(cbox)

		msg := fmt.Sprintf("\nTag '%s' successfully deleted from space '%s'!", selector.Item, selector.Space)
		fmt.Println(aurora.Green(msg))

	},
}

func init() {
	tagsCmd.AddCommand(tagsDeleteCmd)
}
