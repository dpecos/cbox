package tools

import (
	"strings"

	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools/console"
)

func ConsoleReadCommand() *models.Command {
	command := models.Command{
		ID:          console.ReadString("ID"),
		Description: console.ReadString("Description"),
		Details:     console.ReadStringMulti("Details"),
		URL:         console.ReadString("URL"),
		Code:        console.ReadStringMulti("Code / Command"),
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
	command.ID = strings.ToLower(console.EditString("ID", command.ID))
	command.Description = console.EditString("Description", command.Description)
	command.Details = console.EditStringMulti("Details", command.Details)
	command.URL = console.EditString("URL", command.URL)
	command.Code = console.EditStringMulti("Code / Command", command.Code)
}

func ConsoleReadSpace() *models.Space {
	space := models.Space{
		Name:        strings.ToLower(console.ReadString("Name")),
		Description: console.ReadString("Description"),
		Entries:     []models.Command{},
	}
	return &space
}

func ConsoleEditSpace(space *models.Space) {
	space.Name = strings.ToLower(console.EditString("Name", space.Name))
	space.Description = console.EditString("Description", space.Description)
}
