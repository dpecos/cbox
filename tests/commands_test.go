package tests

import (
	"testing"

	"github.com/dpecos/cbox/models"
	uuid "github.com/satori/go.uuid"
)

func createCommand(t *testing.T, space *models.Space) *models.Command {
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

func TestCommandCreationDeletion(t *testing.T) {
	space := createSpace(t)
	command := createCommand(t, space)

	reloadCBox()

	s, _ := cbox.SpaceFind(space.Label)

	if len(space.Entries) != len(s.Entries) {
		t.Fatal("space after persistance has different amount of commands")
	}

	c, err := s.CommandFind(command.ID.String())
	if c == nil || err != nil {
		t.Errorf("could not find command by ID: %v", err)
	}

	c, err = s.CommandFind(command.Label)
	if c == nil || err != nil {
		t.Errorf("could not find command by label: %v", err)
	}

	s.CommandDelete(c)

	reloadCBox()

	_, err = s.CommandFind(command.ID.String())
	if err == nil {
		if err == nil {
			t.Error("command still found after deleting it")
		}
	}
}
