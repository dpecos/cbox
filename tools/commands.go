package tools

import (
	"fmt"
	"strings"

	"github.com/dpecos/cmdbox/models"
	"github.com/logrusorgru/aurora"
)

func PrintCommand(cmd models.Cmd, full bool, sourceOnly bool) {
	if !sourceOnly {
		if len(cmd.Tags) != 0 {
			tags := strings.Join(cmd.Tags, ", ")
			fmt.Printf("%d - (%s) %s - %s\n", aurora.Red(aurora.Bold(cmd.ID)), aurora.Brown(tags), aurora.Blue(aurora.Bold(cmd.Title)), aurora.Green(cmd.CreatedAt))
		} else {
			fmt.Printf("%d - %s - %s\n", aurora.Red(aurora.Bold(cmd.ID)), aurora.Blue(aurora.Bold(cmd.Title)), aurora.Green(cmd.CreatedAt))

		}
		if full {
			if cmd.Description != "" {
				fmt.Printf("\n%s\n", aurora.Cyan(cmd.Description))
			}
			if cmd.URL != "" {
				fmt.Printf("\n%s\n", aurora.Blue(cmd.URL))
			}
			fmt.Printf("\n%s\n\n", cmd.Cmd)
		}
	} else {
		fmt.Println(cmd.Cmd)
	}
}
