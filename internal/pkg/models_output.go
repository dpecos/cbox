package pkg

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dplabs/cbox/internal/pkg/console"
	"github.com/dplabs/cbox/pkg/models"
)

var (
	// spaceIDColor    = console.ColorBoldBlack
	spaceLabelColor = console.ColorBoldGreen
	// idColor          = console.ColorBoldBlack
	labelColor       = console.ColorBoldBlue
	tagsColor        = console.ColorRed
	descriptionColor = fmt.Sprintf
	dateColor        = console.ColorBoldBlack
	urlColor         = console.ColorGreen
	separatorColor   = console.ColorYellow
	starColor        = console.ColorBoldBlack
)

func PrintCommand(header string, cmd *models.Command, full bool, sourceOnly bool) {

	if sourceOnly {
		fmt.Println(cmd.Code)
	} else {

		printHeader(header)

		if !full {
			fmt.Printf(starColor("* "))
		}

		if len(cmd.Tags) != 0 {
			tags := strings.Join(cmd.Tags, ", ")
			fmt.Printf("%s - %s (%s) %s\n", labelColor(cmd.Label), descriptionColor(cmd.Description), tagsColor(tags), dateColor(cmd.CreatedAt.String()))
		} else {
			fmt.Printf("%s - %s %s\n", labelColor(cmd.Label), descriptionColor(cmd.Description), dateColor(cmd.CreatedAt.String()))
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

	sortCommands(commands)

	for i, command := range commands {
		PrintCommand("", command, full, sourceOnly)
		if full && i != len(commands)-1 {
			fmt.Printf("\n%s\n\n", separatorColor("- - - - - - - - - - - -"))
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
	fmt.Printf("%s - %s %s\n", spaceLabelColor(space.Label), descriptionColor(space.Description), dateColor(timestamp))
	printFooter(header)
}

func PrintSelector(header string, selector *models.Selector) {
	printHeader(header)
	fmt.Printf("%s\n", spaceLabelColor(selector.String()))
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
		return strings.Compare(commands[i].Label, commands[j].Label) == -1
	})
}
