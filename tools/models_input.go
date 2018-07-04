package tools

import (
	"strings"

	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools/console"
)

func ConsoleReadCommand() *models.Command {
	command := models.Command{
		ID:          strings.ToLower(console.ReadString("ID", console.NOT_EMPTY_VALUES)),
		Description: console.ReadString("Description"),
		Details:     console.ReadString("Details", console.MULTILINE),
		URL:         console.ReadString("URL"),
		Code:        console.ReadString("Code / Command", console.MULTILINE, console.NOT_EMPTY_VALUES),
		Tags:        []string{},
	}
	tags := console.ReadString("Tags (separated by space)")
	for _, tag := range strings.Split(tags, " ") {
		if tag != "" {
			command.Tags = append(command.Tags, tag)
		}
	}

	return &command
}

func ConsoleEditCommand(command *models.Command) {
	command.ID = strings.ToLower(console.EditString("ID", command.ID, console.NOT_EMPTY_VALUES))
	command.Description = console.EditString("Description", command.Description)
	command.Details = console.EditString("Details", command.Details, console.MULTILINE)
	command.URL = console.EditString("URL", command.URL)
	command.Code = console.EditString("Code / Command", command.Code, console.MULTILINE, console.NOT_EMPTY_VALUES)
}

func ConsoleReadSpace() *models.Space {
	space := models.Space{
		Name:        strings.ToLower(console.ReadString("Name", console.NOT_EMPTY_VALUES)),
		Description: console.ReadString("Description"),
		Entries:     []models.Command{},
	}
	return &space
}

func ConsoleEditSpace(space *models.Space) {
	space.Name = strings.ToLower(console.EditString("Name", space.Name, console.NOT_EMPTY_VALUES))
	space.Description = console.EditString("Description", space.Description)
}
