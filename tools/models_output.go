package tools

import (
	"fmt"
	"strings"
	"time"

	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools/console"
)

var (
	spaceIDColor     = console.ColorBoldBlack
	spaceLabelColor  = console.ColorBlue
	idColor          = console.ColorBoldBlack
	labelColor       = console.ColorBlue
	tagsColor        = console.ColorRed
	descriptionColor = fmt.Sprintf
	dateColor        = console.ColorBoldBlack
	detailsColor     = console.ColorCyan
	urlColor         = console.ColorGreen
	separatorColor   = console.ColorYellow
)

func PrintCommand(header string, cmd *models.Command, full bool, sourceOnly bool) {

	if sourceOnly {
		fmt.Println(cmd.Code)
	} else {

		if header != "" {
			fmt.Printf(separatorColor("- - - %s - - -\n"), header)
		}

		if !full {
			fmt.Printf("* ")
		}
		t := cmd.CreatedAt.UTC().In(time.Local)
		if len(cmd.Tags) != 0 {
			tags := strings.Join(cmd.Tags, ", ")
			fmt.Printf("%s - %s - %s (%s) %s\n", idColor(cmd.ID.String()), labelColor(cmd.Label), descriptionColor(cmd.Description), tagsColor(tags), dateColor(DateToString(t)))
		} else {
			fmt.Printf("%s - %s - %s %s\n", idColor(cmd.ID.String()), labelColor(cmd.Label), descriptionColor(cmd.Description), dateColor(DateToString(t)))
		}
		if full {
			if cmd.Details != "" {
				fmt.Printf("\n%s\n", detailsColor(cmd.Details))
			}
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

func PrintCommandList(commands []models.Command, full bool, sourceOnly bool) {
	for i, command := range commands {
		PrintCommand("", &command, full, sourceOnly)
		if full && i != len(commands)-1 {
			fmt.Printf("\n%s\n\n", separatorColor("- - - - - - - - - - - -"))
		}
	}
}

func PrintTag(tag string) {
	fmt.Printf("%s\n", tagsColor(tag))
}

func PrintSpace(header string, space *models.Space) {
	if header != "" {
		fmt.Printf(separatorColor("- - - %s - - -\n"), header)
	}

	fmt.Printf("%s - %s - %s\n", spaceIDColor(space.ID.String()), spaceLabelColor(space.Label), descriptionColor(space.Description))

	if header != "" {
		fmt.Printf("%s\n\n", separatorColor("- - - - - - - - - - - -"))
	}
}

func PrintSetting(config string, value string) {
	fmt.Printf("%s -> %s\n", console.ColorGreen(config), console.ColorYellow(value))
}
