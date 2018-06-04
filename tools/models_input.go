package tools

import (
	"strings"
	"time"

	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools/console"
)

func ConsoleReadCommand() *models.Command {

	command := models.Command{
		ID:          console.ReadString("ID"),
		Title:       console.ReadString("Title"),
		Description: console.ReadStringMulti("Description"),
		URL:         console.ReadString("URL"),
		Code:        console.ReadStringMulti("Code / Command"),
		Tags:        []string{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tags := console.ReadString("Tags (separated by space)")
	for _, tag := range strings.Split(tags, " ") {
		if tag != "" {
			command.Tags = append(command.Tags, tag)
		}
	}

	return &command
}
