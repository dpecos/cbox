package console

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools/tty"
)

var (
	spaceColor                 = tty.ColorBoldGreen
	atSeparatorColor           = tty.ColorBoldRed
	namespaceSeparatorColor    = tty.ColorBoldWhite
	namespaceColorUser         = tty.ColorCyan
	namespaceColorOrganization = tty.ColorYellow
	labelColor                 = tty.ColorBoldBlue
	tagsColor                  = tty.ColorRed
	descriptionColor           = fmt.Sprintf
	dateColor                  = tty.ColorBoldBlack
	urlColor                   = tty.ColorGreen
	separatorColor             = tty.ColorYellow
	starColor                  = tty.ColorBoldBlack
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
		parts = append(parts, atSeparatorColor("@"))

		if selector.NamespaceType == models.TypeNone {
			format = format + "%s"
			parts = append(parts, spaceColor(selector.Space))
		} else if selector.NamespaceType == models.TypeUser {
			format = format + "%s%s%s"
			parts = append(parts, namespaceColorUser(selector.Namespace), namespaceSeparatorColor(":"), spaceColor(selector.Space))
		} else {
			format = format + "%s%s%s"
			parts = append(parts, namespaceColorOrganization(selector.Namespace), namespaceSeparatorColor("/"), spaceColor(selector.Space))
		}
	}
	return fmt.Sprintf(format, parts...)
}

func commandSummary(cmd *models.Command) string {
	if len(cmd.Tags) != 0 {
		tags := strings.Join(cmd.Tags, ", ")
		return fmt.Sprintf("%s - %s (%s) %s", selector(cmd.Selector), descriptionColor(cmd.Description), tagsColor(tags), dateColor(cmd.CreatedAt.String()))
	} else {
		return fmt.Sprintf("%s - %s %s", selector(cmd.Selector), descriptionColor(cmd.Description), dateColor(cmd.CreatedAt.String()))
	}
}

func PrintCommand(header string, cmd *models.Command, sourceOnly bool) {

	if cmd == nil {
		log.Fatal("Trying to display a nil command")
	}

	if sourceOnly {
		tty.Print(cmd.Code + "\n")
	} else {

		printHeader(header)
		tty.Print("\n")

		if cmd.Selector.NamespaceType == models.TypeOrganization {
			tty.Print("  Namespace: %s (Organization)\n", namespaceColorOrganization(cmd.Selector.Namespace))
		} else if cmd.Selector.NamespaceType == models.TypeUser {
			tty.Print("  Namespace: %s (User)\n", namespaceColorUser(cmd.Selector.Namespace))
		} else {
			tty.Print("  Namespace: -\n")
		}
		tty.Print("  Space: %s \n", spaceColor(cmd.Selector.Space))
		tty.Print("  Label: %s \n", labelColor(cmd.Label))
		tty.Print("  Selector: %s \n", selector(cmd.Selector))
		tty.Print("\n")
		tty.Print("  Description: %s\n", descriptionColor(cmd.Description))
		tty.Print("  URL: %s\n", urlColor(cmd.URL))
		tty.Print("  Tags: %s\n", tagsColor(strings.Join(cmd.Tags, ", ")))
		tty.Print("\n")
		tty.Print("  Created at: %s\n", dateColor(cmd.CreatedAt.String()))
		tty.Print("  Updated at: %s\n", dateColor(cmd.UpdatedAt.String()))

		tty.Print("\n%s\n\n%s\n\n", separatorColor("- - -"), cmd.Code)

		printFooter(header)
	}
}

func runFZFList(header string, commands []*models.Command) {
	args := []string{"--ansi", "--exact", "--preview-window=down:30%:wrap", "--preview", "echo {} | cut -f1 -d' ' | xargs cbox command view"}
	if header != "" {
		args = append(args, "--header="+header)
	}
	fzfProcess := exec.Command("fzf", args...)
	stdin, err := fzfProcess.StdinPipe()
	if err != nil {
		log.Fatalf("console: interactive mode: failed to spawn process and get its stdin: %v", err)
	}

	fzfProcess.Stderr = os.Stderr

	for _, cmd := range commands {
		io.WriteString(stdin, commandSummary(cmd))
		io.WriteString(stdin, "\n")
	}

	stdin.Close()

	out, err := fzfProcess.Output()
	if err != nil {
		exitCode := fzfProcess.ProcessState.ExitCode()
		if exitCode == -1 {
			log.Fatalf("console: interactive mode: 'fzf' failed to start: %v", err)
		} else if exitCode == 2 {
			log.Fatalf("console: interactive mode: 'fzf' returned an internal error: %v", err)
		}
		return
	}

	selector := strings.Split(string(out), " ")[0]

	for _, cmd := range commands {
		if cmd.ID == selector {
			PrintCommand(selector, cmd, false)
			break
		}
	}
}

func staticCommandList(header string, commands []*models.Command) {
	printHeader(header)

	if len(commands) != 0 {
		sortCommands(commands)

		for _, command := range commands {
			tty.Print(" * %s\n", commandSummary(command))
		}
	}

	printFooter(header)
}

func PrintCommandList(header string, commands []*models.Command, interactive bool) {
	if interactive {
		runFZFList(header, commands)
	} else {
		staticCommandList(header, commands)
	}
}

func PrintTag(tag string) {
	tty.Print("%s %s\n", starColor("*"), tagsColor(tag))
}

func PrintSpace(header string, space *models.Space) {
	printHeader(header)
	timestamp := fmt.Sprintf("(Last updated: %s - Created: %s)", space.UpdatedAt.String(), space.CreatedAt.String())
	tty.Print("%s - %s %s\n", selector(space.Selector), descriptionColor(space.Description), dateColor(timestamp))
	printFooter(header)
}

func PrintSelector(header string, s *models.Selector) {
	printHeader(header)
	tty.Print("%s\n", selector(s))
	printFooter(header)
}

func PrintSetting(config string, value string) {
	tty.Print("%s -> %s\n", tty.ColorGreen(config), tty.ColorYellow(value))
}

func printHeader(header string) {
	if header != "" {
		tty.Print(separatorColor("- - - %s - - -\n"), header)
	}
}

func printFooter(header string) {
	if header != "" {
		tty.Print("%s\n\n", separatorColor("- - - - - - - - - - - -"))
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
