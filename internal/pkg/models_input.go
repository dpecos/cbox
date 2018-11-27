package pkg

import (
	"fmt"
	"log"
	"strings"

	"github.com/dplabs/cbox/internal/pkg/console"
	"github.com/dplabs/cbox/pkg/models"
	"github.com/gofrs/uuid"
)

func ConsoleReadCommand() *models.Command {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("could not generate a new ID: %v", err)
	}
	command := models.Command{
		Label:       strings.ToLower(console.ReadString("Label", console.ONLY_VALID_CHARS, console.NOT_EMPTY_VALUES)),
		Description: console.ReadString("Description"),
		URL:         console.ReadString("URL"),
		Code:        console.ReadString("Code / Command", console.MULTILINE, console.NOT_EMPTY_VALUES),
		Tags:        []string{},
	}
	command.ID = id

	tags := console.ReadString("Tags (separated by space)")
	for _, tag := range strings.Split(tags, " ") {
		if tag != "" {
			command.Tags = append(command.Tags, tag)
		}
	}

	fmt.Println()

	return &command
}

func ConsoleEditCommand(command *models.Command) {
	command.Label = strings.ToLower(console.EditString("Label", command.Label, console.ONLY_VALID_CHARS, console.NOT_EMPTY_VALUES))
	command.Description = console.EditString("Description", command.Description)
	command.URL = console.EditString("URL", command.URL)
	command.Code = console.EditString("Code / Command", command.Code, console.MULTILINE, console.NOT_EMPTY_VALUES)

	fmt.Println()
}

func ConsoleReadSpace() *models.Space {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("could not generate a new ID: %v", err)
	}
	space := models.Space{
		Label:       strings.ToLower(console.ReadString("Label", console.NOT_EMPTY_VALUES, console.ONLY_VALID_CHARS)),
		Description: console.ReadString("Description"),
		Entries:     []models.Command{},
	}
	space.ID = id

	fmt.Println()

	return &space
}

func ConsoleEditSpace(space *models.Space) {
	space.Label = strings.ToLower(console.EditString("Label", space.Label, console.NOT_EMPTY_VALUES, console.ONLY_VALID_CHARS))
	space.Description = console.EditString("Description", space.Description)

	fmt.Println()
}
