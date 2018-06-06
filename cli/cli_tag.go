package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:     "tag",
	Aliases: []string{"t"},
	Args:    cobra.MinimumNArgs(2),
	Short:   "Add tags to a command",
	Long:    tools.Logo,
	Run: func(cmd *cobra.Command, args []string) {

		selector, err := models.ParseSelector(args[0])
		if err != nil {
			log.Fatal("Could not parse selector", err)
		}

		cbox := core.LoadCbox()

		space := cbox.SpaceFind(selector.Space)
		command := space.CommandFind(selector.Item)

		for _, tag := range args[1:] {
			command.TagAdd(strings.ToLower(tag))
		}

		core.PersistCbox(cbox)

		tools.PrintCommand(command, true, false)
		fmt.Println(aurora.Green("\nCommand tagged successfully!\n"))
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
}
