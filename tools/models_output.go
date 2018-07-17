package tools

import (
	"fmt"
	"strings"
	"time"

	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools/console"
)

var (
	spaceIDColor     = console.ColorBlue
	idColor          = console.ColorBlue
	tagsColor        = console.ColorRed
	descriptionColor = fmt.Sprintf
	dateColor        = console.ColorBoldBlack
	detailsColor     = console.ColorCyan
	urlColor         = console.ColorGreen
)

func PrintCommand(cmd *models.Command, full bool, sourceOnly bool) {

	if sourceOnly {
		fmt.Println(cmd.Code)
	} else {
		if !full {
			fmt.Printf("* ")
		}
		t := cmd.CreatedAt.UTC().In(time.Local)
		if len(cmd.Tags) != 0 {
			tags := strings.Join(cmd.Tags, ", ")
			fmt.Printf("%s - %s (%s) %s\n", idColor(cmd.ID), descriptionColor(cmd.Description), tagsColor(tags), dateColor(DateToString(t)))
		} else {
			fmt.Printf("%s - %s %s\n", idColor(cmd.ID), descriptionColor(cmd.Description), dateColor(DateToString(t)))
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
	}
}

func PrintTag(tag string) {
	fmt.Printf("%s\n", tagsColor(tag))
}

func PrintSpace(space *models.Space) {
	fmt.Printf("%s - %s\n", spaceIDColor(space.ID), descriptionColor(space.Description))
}
