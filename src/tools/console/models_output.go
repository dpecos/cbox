package console

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools/tty"
)

var (
	spaceColor              = tty.ColorBoldGreen
	spaceSeparatorColor     = tty.ColorBoldRed
	namespaceSeparatorColor = tty.ColorBoldWhite
	labelColor              = tty.ColorBoldBlue
	tagsColor               = tty.ColorRed
	descriptionColor        = fmt.Sprintf
	dateColor               = tty.ColorBoldBlack
	urlColor                = tty.ColorGreen
	separatorColor          = tty.ColorYellow
	starColor               = tty.ColorBoldBlack
)

func selector(selector *models.Selector) string {
	format := ""
	parts := []interface{}{}

	if selector.Item != "" {
		format = "%s"
		parts = append(parts, labelColor(selector.Item))
	}

	if selector.Space != "" {
		format = format + "%s"
		parts = append(parts, spaceSeparatorColor("@"))

		if selector.NamespaceType == models.TypeNone {
			format = format + "%s"
			parts = append(parts, spaceColor(selector.Space))
		} else if selector.NamespaceType == models.TypeUser {
			format = format + "%s%s%s"
			parts = append(parts, spaceColor(selector.Namespace), namespaceSeparatorColor(":"), spaceColor(selector.Space))
		} else {
			format = format + "%s%s%s"
			parts = append(parts, spaceColor(selector.Namespace), namespaceSeparatorColor("/"), spaceColor(selector.Space))
		}
	}
	return fmt.Sprintf(format, parts...)
}

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

		if len(cmd.Tags) != 0 {
			tags := strings.Join(cmd.Tags, ", ")
			fmt.Printf("%s - %s (%s) %s\n", selector(cmd.Selector), descriptionColor(cmd.Description), tagsColor(tags), dateColor(cmd.CreatedAt.String()))
		} else {
			fmt.Printf("%s - %s %s\n", selector(cmd.Selector), descriptionColor(cmd.Description), dateColor(cmd.CreatedAt.String()))
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
	fmt.Printf("%s - %s %s\n", selector(space.Selector), descriptionColor(space.Description), dateColor(timestamp))
	printFooter(header)
}

func PrintSelector(header string, s *models.Selector) {
	printHeader(header)
	fmt.Printf("%s\n", selector(s))
	printFooter(header)
}

func PrintSetting(config string, value string) {
	fmt.Printf("%s -> %s\n", tty.ColorGreen(config), tty.ColorYellow(value))
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
