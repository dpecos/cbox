package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools"
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

	var spaces []*models.Space = []*models.Space{}
	if sel != "" {
		space, err := findSpace(selector)
		if err != nil {
			log.Fatalf("search: %v", err)
		}
		spaces = append(spaces, space)
	} else {
		spaces = cboxInstance.Spaces
	}

	var commands []*models.Command = []*models.Command{}
	for _, space := range spaces {
		var err error
		cs, err := space.SearchCommands(selector.Item, criteria)
		if err != nil {
			log.Fatalf("search: %v", err)
		}
		commands = append(commands, cs...)
	}

	if selector.Item != "" {
		tools.PrintCommandList(fmt.Sprintf("Results for \"%s\" (within tag: %s)", criteria, selector.Item), commands, showCommandsSource, false)
	} else {
		tools.PrintCommandList(fmt.Sprintf("Results for \"%s\"", criteria), commands, showCommandsSource, false)
	}

}
