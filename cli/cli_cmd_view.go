package cli

import (
	"log"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/cobra"
)

var sourceOnly bool

var viewCmd = &cobra.Command{
	Use:     "view",
	Aliases: []string{"v"},
	Args:    cobra.ExactArgs(1),
	Short:   "View one command",
	Run: func(cmd *cobra.Command, args []string) {

		selector, err := models.ParseSelector(args[0])
		if err != nil {
			log.Fatal("Could not parse selector", err)
		}

		cbox := core.LoadCbox()
		space := cbox.SpaceFind(selector.Space)
		command := space.CommandFind(selector.Item)

		tools.PrintCommand(command, true, sourceOnly)
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
	viewCmd.Flags().BoolVarP(&sourceOnly, "src", "s", false, "view only code snippet source code")
}
