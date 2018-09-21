package tools

import (
	"fmt"
	"strings"

	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools/console"
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

		if header != "" {
			fmt.Printf(separatorColor("- - - %s - - -\n"), header)
		}

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

		if header != "" {
			fmt.Printf("%s\n\n", separatorColor("- - - - - - - - - - - -"))
		}
	}
}

func PrintCommandList(header string, commands []models.Command, full bool, sourceOnly bool) {
	if header != "" {
		fmt.Printf(separatorColor("- - - %s - - -\n"), header)
	}

	for i, command := range commands {
		PrintCommand("", &command, full, sourceOnly)
		if full && i != len(commands)-1 {
			fmt.Printf("\n%s\n\n", separatorColor("- - - - - - - - - - - -"))
		}
	}

	if header != "" {
		fmt.Printf("%s\n\n", separatorColor("- - - - - - - - - - - -"))
	}
}

func PrintTag(tag string) {
	fmt.Printf("%s %s\n", starColor("*"), tagsColor(tag))
}

func PrintSpace(header string, space *models.Space) {
	if header != "" {
		fmt.Printf(separatorColor("- - - %s - - -\n"), header)
	}

	fmt.Printf("%s - %s %s\n", spaceLabelColor(space.Label), descriptionColor(space.Description), dateColor(space.CreatedAt.String()))

	if header != "" {
		fmt.Printf("%s\n\n", separatorColor("- - - - - - - - - - - -"))
	}
}

func PrintSetting(config string, value string) {
	fmt.Printf("%s -> %s\n", console.ColorGreen(config), console.ColorYellow(value))
}
