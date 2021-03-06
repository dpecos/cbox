package integration_tests

import (
	"testing"

	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/tests"
)

func createCommand(t *testing.T, space *models.Space) *models.Command {
	command := models.Command{
		Label:       tests.RandString(8),
		Description: tests.RandString(15),
		URL:         tests.RandString(15),
		Code:        tests.RandString(30),
		Tags:        []string{"test"},
	}
	command.Selector = space.Selector.CloneForItem(command.Label)

	space.CommandAdd(&command, false)

	return &command
}

func TestCommandCreationDeletion(t *testing.T) {
	cboxInstance := tests.InitializeCBox()

	space := tests.CreateSpace(t, cboxInstance)
	command := createCommand(t, space)

	cboxInstance = tests.ReloadCBox(cboxInstance)

	s, _ := cboxInstance.SpaceFind(space.Selector.NamespaceType, space.Selector.Namespace, space.Label)

	if s == nil {
		t.Fatal("could not find space")
	}

	if len(space.Entries) != len(s.Entries) {
		t.Fatal("space after persistence has different amount of commands")
	}

	c, err := s.CommandFind(command.Label)
	if c == nil || err != nil {
		t.Errorf("could not find command by label: %v", err)
	}

	s.CommandDelete(c)

	cboxInstance = tests.ReloadCBox(cboxInstance)

	_, err = s.CommandFind(command.Label)
	if err == nil {
		if err == nil {
			t.Error("command still found after deleting it")
		}
	}
}

func TestCommandEdition(t *testing.T) {
	cboxInstance := tests.InitializeCBox()

	space := tests.CreateSpace(t, cboxInstance)
	command := createCommand(t, space)

	cboxInstance = tests.ReloadCBox(cboxInstance)

	s, err := cboxInstance.SpaceFind(space.Selector.NamespaceType, space.Selector.Namespace, space.Label)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.CommandFind(command.Label)
	if err != nil {
		t.Fatal(err)
	}

	previousLabel := c.Label
	newLabel := tests.RandString(8)

	c.Label = newLabel

	space.CommandEdit(c, previousLabel)

	cboxInstance = tests.ReloadCBox(cboxInstance)

	s, _ = cboxInstance.SpaceFind(space.Selector.NamespaceType, space.Selector.Namespace, space.Label)

	c, err = s.CommandFind(previousLabel)
	if err == nil {
		t.Fatalf("command found using old label")
	}

	c, err = s.CommandFind(newLabel)
	if err != nil {
		t.Fatalf("command not found using new label")
	}
}

func TestCommandLabelUniquenessOnCreation(t *testing.T) {
	cboxInstance := tests.InitializeCBox()

	space := tests.CreateSpace(t, cboxInstance)
	c1 := createCommand(t, space)

	c2 := models.Command{
		Label:       c1.Label,
		Description: tests.RandString(15),
		URL:         tests.RandString(15),
		Code:        tests.RandString(30),
		Tags:        []string{"test"},
	}

	err := space.CommandAdd(&c2, false)
	if err == nil {
		t.Fatalf("labels have to be unique within an space after adding a command")
	}
}

func TestCommandLabelUniquenessOnEdition(t *testing.T) {
	cboxInstance := tests.InitializeCBox()

	space := tests.CreateSpace(t, cboxInstance)
	c1 := createCommand(t, space)
	c2 := createCommand(t, space)

	previousLabel := c2.Label
	c2.Label = c1.Label

	err := space.CommandEdit(c2, previousLabel)
	if err == nil {
		t.Fatalf("labels have to be unique within an space after editing a command")
	}
}

func TestCommandTagging(t *testing.T) {
	cboxInstance := tests.InitializeCBox()

	space := tests.CreateSpace(t, cboxInstance)
	c1 := createCommand(t, space)

	c1.TagAdd("tag-ok")

	if !tests.AssertSliceContains(c1.Tags, "tag-ok") {
		t.Errorf("Added tag not found")
	}

	c1.TagAdd("tag-ok")

	if !tests.AssertSliceContains(c1.Tags, "tag-ok") {
		t.Errorf("Removed tag still found")
	}
}
