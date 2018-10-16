package tests

import (
	"sort"
	"testing"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
	"github.com/gofrs/uuid"
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

	err = cbox.SpaceDestroy(space)
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

	err := cbox.SpaceCreate(&s2)
	if err == nil {
		t.Fatalf("space labels have to be unique")
	}
}

func TestDeleteSpace(t *testing.T) {
	setupTests()

	s1 := createSpace(t)
	s2 := createSpace(t)
	s3 := createSpace(t)

	reloadCBox()

	expected := []string{"default", s1.Label, s3.Label}

	cbox.SpaceDestroy(s2)
	core.SpaceDeleteFile(s2)

	reloadCBox()

	result := cbox.SpaceLabels()

	sort.Strings(expected)
	sort.Strings(result)

	if !assertSliceEqual(expected, result) {
		t.Errorf("space deletion did return a different result from expected: %v - %v", expected, result)
	}
}

func TestSearch(t *testing.T) {
	setupTests()

	s1 := createSpace(t)
	c1 := createCommand(t, s1)

	reloadCBox()

	criteria := c1.Label[0:3]
	result, err := s1.SearchCommands(criteria)

	if err != nil {
		t.Errorf("space search errored: %s", err)
	}

	if len(result) != 1 || result[0].ID != c1.ID {
		t.Errorf("space search did not return the expected command: %v", c1)
	}
}
