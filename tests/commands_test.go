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

func TestCommandEdition(t *testing.T) {
	space := createSpace(t)
	command := createCommand(t, space)

	reloadCBox()

	s, _ := cbox.SpaceFind(space.Label)
	c, _ := s.CommandFind(command.ID.String())

	previousLabel := c.Label
	newLabel := randString(8)

	c.Label = newLabel

	space.CommandEdit(c, previousLabel)

	reloadCBox()

	s, _ = cbox.SpaceFind(space.Label)

	c, err := s.CommandFind(previousLabel)
	if err == nil {
		t.Fatalf("command found using old label")
	}

	c, err = s.CommandFind(newLabel)
	if err != nil {
		t.Fatalf("command not found using new label")
	}
}

func TestCommandLabelUniquenessOnCreation(t *testing.T) {
	space := createSpace(t)
	c1 := createCommand(t, space)

	id, _ := uuid.NewV4()
	c2 := models.Command{
		ID:          id,
		Label:       c1.Label,
		Description: randString(15),
		Details:     randString(15),
		URL:         randString(15),
		Code:        randString(30),
		Tags:        []string{"test"},
	}

	err := space.CommandAdd(&c2)
	if err == nil {
		t.Fatalf("labels have to be unique within an space after adding a command")
	}
}

func TestCommandLabelUniquenessOnEdition(t *testing.T) {
	space := createSpace(t)
	c1 := createCommand(t, space)
	c2 := createCommand(t, space)

	previousLabel := c2.Label
	c2.Label = c1.Label

	err := space.CommandEdit(c2, previousLabel)
	if err == nil {
		t.Fatalf("labels have to be unique within an space after editing a command")
	}
}
