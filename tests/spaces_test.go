package tests

import (
	"testing"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
	uuid "github.com/satori/go.uuid"
)

func TestSpaceCreationDeletion(t *testing.T) {
	space := createSpace(t)

	reloadCBox()

	assertSpaceFileExists(t, space)

	s, err := cbox.SpaceFind(space.ID.String())
	if s == nil || err != nil {
		t.Errorf("could not find space by ID: %v", err)
	}

	s, err = cbox.SpaceFind(space.Label)
	if s == nil || err != nil {
		t.Errorf("could not find space by label: %v", err)
	}

	err = cbox.SpaceDelete(space)
	if err != nil {
		t.Error(err)
	}
	core.SpaceDeleteFile(space)

	assertSpaceFileNotExists(t, space)

	reloadCBox()

	_, err = cbox.SpaceFind(space.ID.String())
	if err == nil {
		t.Error("space still found after deleting it")
	}

}

func TestSpaceLabelUniquenessOnCreation(t *testing.T) {
	s1 := createSpace(t)

	id, _ := uuid.NewV4()
	s2 := models.Space{
		Label:       s1.Label,
		Description: randString(15),
	}
	s2.ID = id

	err := cbox.SpaceAdd(&s2)
	if err == nil {
		t.Fatalf("space labels have to be unique")
	}
}
