package tests

import (
	"testing"

	"github.com/dpecos/cbox/models"
	uuid "github.com/satori/go.uuid"
)

func createCommand(t *testing.T) *models.Command {
	id, _ := uuid.NewV4()

	command := models.Command{
		ID:          id,
		Label:       randString(8),
		Description: randString(15),
		Details:     randString(15),
		URL:         randString(15),
		Code:        randString(30),
		Tags:        []string{"test"},
	}

	space.CommandAdd(&command)

	return &command
}

func TestCommandCreation(t *testing.T) {
	space = createSpace(t)

	// command := createCommand(t)

	// reloadCBox()

	// s := cbox.SpaceFind(space.Label)
	// c := s.CommandFind(command.Label)

	// if c == nil {
	// 	t.Error("could not find command by label")
	// }
}
