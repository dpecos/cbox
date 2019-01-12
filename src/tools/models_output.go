package tools

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools/console"
)

var (
	spaceColor          = console.ColorBoldGreen
	spaceSeparatorColor = console.ColorBoldRed
	labelColor          = console.ColorBoldBlue
	tagsColor           = console.ColorRed
	descriptionColor    = fmt.Sprintf
	dateColor           = console.ColorBoldBlack
	urlColor            = console.ColorGreen
	separatorColor      = console.ColorYellow
	starColor           = console.ColorBoldBlack
)

func PrintCommand(header string, cmd *models.Command, full bool, sourceOnly bool) {

	if cmd == nil {
		log.Fatal("Trying to display a nil command")
	}

	if sourceOnly {
		fmt.Println(cmd.Code)
	} else {

		printHeader(header)

		if !full {
			fmt.Printf(starColor("* "))
		}

		cmdStr := fmt.Sprintf("%s%s%s", labelColor(cmd.Label), spaceSeparatorColor("@"), spaceColor(cmd.Space.String()))

		if len(cmd.Tags) != 0 {
			tags := strings.Join(cmd.Tags, ", ")
			fmt.Printf("%s - %s (%s) %s\n", cmdStr, descriptionColor(cmd.Description), tagsColor(tags), dateColor(cmd.CreatedAt.String()))
		} else {
			fmt.Printf("%s - %s %s\n", cmdStr, descriptionColor(cmd.Description), dateColor(cmd.CreatedAt.String()))
		}

		if full {
			if cmd.URL != "" {
				fmt.Printf("\n%s\n", urlColor(cmd.URL))
			}
			fmt.Printf("\n%s\n", cmd.Code)
		}

		printFooter(header)
	}
}

func PrintCommandList(header string, commands []*models.Command, full bool, sourceOnly bool) {
	printHeader(header)

	if len(commands) != 0 {
		sortCommands(commands)

		for i, command := range commands {
			PrintCommand("", command, full, sourceOnly)
			if full && i != len(commands)-1 {
				fmt.Printf("\n%s\n\n", separatorColor("- - - - - - - - - - - -"))
			}
		}
	}

	printFooter(header)
}

func PrintTag(tag string) {
	fmt.Printf("%s %s\n", starColor("*"), tagsColor(tag))
}

func PrintSpace(header string, space *models.Space) {
	printHeader(header)
	timestamp := fmt.Sprintf("(Last updated: %s - Created: %s)", space.UpdatedAt.String(), space.CreatedAt.String())
	fmt.Printf("%s - %s %s\n", spaceColor(space.String()), descriptionColor(space.Description), dateColor(timestamp))
	printFooter(header)
}

func PrintSelector(header string, selector *models.Selector) {
	printHeader(header)
	fmt.Printf("%s\n", spaceColor(selector.String()))
	printFooter(header)
}

func PrintSetting(config string, value string) {
	fmt.Printf("%s -> %s\n", console.ColorGreen(config), console.ColorYellow(value))
}

func printHeader(header string) {
	if header != "" {
		fmt.Printf(separatorColor("- - - %s - - -\n"), header)
	}
}

func printFooter(header string) {
	if header != "" {
		fmt.Printf("%s\n\n", separatorColor("- - - - - - - - - - - -"))
	}
}

func sortCommands(commands []*models.Command) {
	sort.Slice(commands, func(i, j int) bool {
		if commands[i] == nil || commands[j] == nil {
			log.Fatal("Trying to sort a list of commands with nil entries")
		}
		return strings.Compare(commands[i].Label, commands[j].Label) == -1
	})
}
