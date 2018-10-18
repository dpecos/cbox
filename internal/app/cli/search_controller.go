package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/dpecos/cbox/internal/app/core"
	"github.com/dpecos/cbox/internal/pkg"
	"github.com/spf13/cobra"
)

func (ctrl *CLIController) SearchCommands(cmd *cobra.Command, args []string) {

	var sel, criteria string

	if len(args) == 2 {
		sel = args[0]
		criteria = args[1]
	} else if len(args) == 1 {
		if strings.Contains(args[0], "@") {
			log.Fatalf("search: criteria not specified - this looks like a selector")
		}
		sel = ""
		criteria = args[0]
	} else {
		log.Fatal("search: criteria not specified")
	}

	selector := ctrl.parseSelector([]string{sel})

	cbox := core.LoadCbox("")
	space, err := cbox.SpaceFind(selector.Space)
	if err != nil {
		log.Fatalf("search: %v", err)
	}

	commands, err := space.SearchCommands(selector.Item, criteria)
	if err != nil {
		log.Fatalf("search: %v", err)
	}

	if selector.Item != "" {
		pkg.PrintCommandList(fmt.Sprintf("Results for \"%s\" (within tag: %s)", criteria, selector.Item), commands, viewSnippet, false)
	} else {
		pkg.PrintCommandList(fmt.Sprintf("Results for \"%s\"", criteria), commands, viewSnippet, false)
	}

}
