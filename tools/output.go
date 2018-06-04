package tools

import (
	"fmt"
	"strings"
	"time"

	"github.com/dpecos/cmdbox/models"
	"github.com/logrusorgru/aurora"
)

func PrintCommand(cmd models.Command, full bool, sourceOnly bool) {
	if !sourceOnly {
		if len(cmd.Tags) != 0 {
			tags := strings.Join(cmd.Tags, ", ")
			fmt.Printf("%d - (%s)", aurora.Red(cmd.ID), aurora.Brown(tags))
		} else {
			fmt.Printf("%d -", aurora.Red(cmd.ID))
		}
		t := cmd.CreatedAt.UTC().In(time.Local)
		fmt.Printf(" %s %s\n", aurora.Blue(aurora.Bold(cmd.Title)), aurora.Gray(DateToString(t)))
		if full {
			if cmd.Description != "" {
				fmt.Printf("\n%s\n", aurora.Green(cmd.Description))
			}
			if cmd.URL != "" {
				fmt.Printf("\n%s\n", aurora.Blue(cmd.URL))
			}
			fmt.Printf("\n%s\n\n", cmd.Code)
		}
	} else {
		fmt.Println(cmd.Code)
	}
}

func PrintTag(tag string) {
	fmt.Printf("%s\n", aurora.Brown(tag))
}

func PrintSpace(space models.Space) {
	fmt.Printf("%s - %s\n", aurora.Green(space.Name), aurora.Blue(aurora.Bold(space.Title)))
}
