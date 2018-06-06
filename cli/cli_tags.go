package cli

import (
	"fmt"
	"log"
	"sort"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/cobra"
)

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Args:  cobra.MaximumNArgs(1),
	Short: "List the tags available in your cbox",
	Long:  tools.Logo,
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

		tags := space.TagsList(selector.Item)
		sort.Strings(tags)

		for _, tag := range tags {
			fmt.Printf("%s\n", tag)
		}

	},
}

func init() {
	rootCmd.AddCommand(tagsCmd)
}
