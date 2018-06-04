package cli

import (
	"log"

	"github.com/dpecos/cmdbox/core"
	"github.com/dpecos/cmdbox/models"
	"github.com/dpecos/cmdbox/tools"
	"github.com/spf13/cobra"
)

var (
	viewSnippet bool
	filterTag   string
)
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List the content of your cmdbox",
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
		if space == nil {
			log.Fatalf("Could not find space %s", selector.Space)
		}
		commands := space.CommandList(selector.Item)

		for _, command := range commands {
			tools.PrintCommand(command, viewSnippet, false)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&viewSnippet, "view", "v", false, "Show all details about commands")
}
