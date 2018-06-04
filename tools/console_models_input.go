package tools

import (
	"strings"
	"time"

	"github.com/dpecos/cbox/models"
)

func ConsoleReadCommand() *models.Command {

	command := models.Command{
		ID:          ReadString("ID"),
		Title:       ReadString("Title"),
		Description: ReadStringMulti("Description"),
		URL:         ReadString("URL"),
		Code:        ReadStringMulti("Code / Command"),
		Tags:        []string{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tags := ReadString("Tags (separated by space)")
	for _, tag := range strings.Split(tags, " ") {
		if tag != "" {
			command.Tags = append(command.Tags, tag)
		}
	}

	return &command
}
