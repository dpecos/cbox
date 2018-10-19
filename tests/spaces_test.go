package tests

import (
	"sort"
	"testing"

	"github.com/dpecos/cbox/internal/app/core"
	"github.com/dpecos/cbox/pkg/models"
	"github.com/gofrs/uuid"
)

func TestSpaceCreationDeletion(t *testing.T) {
	cboxInstance := initializeCBox()

	space := createSpace(t, cboxInstance)

	cboxInstance = reloadCBox(cboxInstance)

	assertSpaceFileExists(t, cboxInstance, space)

	s, err := cboxInstance.SpaceFind(space.ID.String())
	if s == nil || err != nil {
		t.Errorf("could not find space by ID: %v", err)
	}

	s, err = cboxInstance.SpaceFind(space.Label)
	if s == nil || err != nil {
		t.Errorf("could not find space by label: %v", err)
	}

	err = cboxInstance.SpaceDestroy(space)
	if err != nil {
		t.Error(err)
	}
	core.DeleteSpaceFile(space)

	assertSpaceFileNotExists(t, cboxInstance, space)

	cboxInstance = reloadCBox(cboxInstance)

	_, err = cboxInstance.SpaceFind(space.ID.String())
	if err == nil {
		t.Error("space still found after deleting it")
	}

}

func TestSpaceLabelUniquenessOnCreation(t *testing.T) {
	cboxInstance := initializeCBox()

	s1 := createSpace(t, cboxInstance)

	id, _ := uuid.NewV4()
	s2 := models.Space{
		Label:       s1.Label,
		Description: randString(15),
	}
	s2.ID = id

	err := cboxInstance.SpaceCreate(&s2)
	if err == nil {
		t.Fatalf("space labels have to be unique")
	}
}

func TestDeleteSpace(t *testing.T) {
	cboxInstance := initializeCBox()

	s1 := createSpace(t, cboxInstance)
	s2 := createSpace(t, cboxInstance)
	s3 := createSpace(t, cboxInstance)

	cboxInstance = reloadCBox(cboxInstance)

	expected := []string{"default", s1.Label, s3.Label}

	cboxInstance.SpaceDestroy(s2)
	core.DeleteSpaceFile(s2)

	cboxInstance = reloadCBox(cboxInstance)

	result := cboxInstance.SpaceLabels()

	sort.Strings(expected)
	sort.Strings(result)

	if !assertSliceEqual(expected, result) {
		t.Errorf("space deletion did return a different result from expected: %v - %v", expected, result)
	}
}

func TestSearch(t *testing.T) {
	cboxInstance := initializeCBox()

	s1 := createSpace(t, cboxInstance)
	c1 := createCommand(t, s1)

	cboxInstance = reloadCBox(cboxInstance)

	criteria := c1.Label[0:3]
	result, err := s1.SearchCommands("", criteria)

	if err != nil {
		t.Errorf("space search error: %v", err)
	}

	if len(result) != 1 || result[0].ID != c1.ID {
		t.Errorf("space search did not return the expected command: %v", c1)
	}
}

func TestSearchWithinTag(t *testing.T) {
	cboxInstance := initializeCBox()

	s1 := createSpace(t, cboxInstance)
	c1 := createCommand(t, s1)

	cboxInstance = reloadCBox(cboxInstance)

	criteria := c1.Label[0:3]
	result, err := s1.SearchCommands("test", criteria)

	if err != nil {
		t.Errorf("space search error: %v", err)
	}

	if len(result) != 1 || result[0].ID != c1.ID {
		t.Errorf("space search did not return the expected command: %v", c1)
	}

	result, err = s1.SearchCommands("non-existing", criteria)

	if err != nil {
		t.Errorf("space search error (non-existing): %v", err)
	}

	if len(result) != 0 {
		t.Errorf("space search did return something for non-existing tag: %v", result)
	}
}
