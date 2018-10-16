package cli

import (
	"fmt"
	"log"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) SearchCommands(cmd *cobra.Command, args []string) {

	var sel, criteria string

	if len(args) == 2 {
		sel = args[0]
		criteria = args[1]
	} else if len(args) == 1 {
		sel = ""
		criteria = args[0]
	} else {
		log.Fatalf("search: inccorrect number of parameters: %d", len(args))
	}

	selector := ctrl.parseSelector([]string{sel})

	cbox := core.LoadCbox("")
	space, err := cbox.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("search: %v", err)
	}

	commands, err := space.SearchCommands(criteria)
	if err != nil {
		log.Fatalf("search: %v", err)
	}

	tools.PrintCommandList(fmt.Sprintf("Results for \"%s\"", criteria), commands, viewSnippet, false)
}
