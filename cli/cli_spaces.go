package cli

import (
	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/cobra"
)

var spacesCmd = &cobra.Command{
	Use:   "spaces",
	Short: "Show available spaces in your cbox",
	Long:  tools.Logo,
	Run: func(cmd *cobra.Command, args []string) {

		spaces := core.SpaceList()
		for _, space := range spaces {
			tools.PrintSpace(space)
		}
	},
}

func init() {
	rootCmd.AddCommand(spacesCmd)
}
