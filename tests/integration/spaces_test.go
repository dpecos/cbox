package integration_tests

import (
	"sort"
	"testing"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
)

func TestSpaceCreationDeletion(t *testing.T) {
	cboxInstance := initializeCBox()

	space := createSpace(t, cboxInstance)

	cboxInstance = reloadCBox(cboxInstance)

	assertSpaceFileExists(t, cboxInstance, space)

	s, err := cboxInstance.SpaceFind(space.Selector.NamespaceType, space.Selector.Namespace, space.Label)
	if s == nil || err != nil {
		t.Errorf("could not find space by label: %v", err)
	}

	err = cboxInstance.SpaceDestroy(space)
	if err != nil {
		t.Error(err)
	}
	core.DeleteSpaceFile(space.Selector)

	assertSpaceFileNotExists(t, cboxInstance, space)

	cboxInstance = reloadCBox(cboxInstance)

	_, err = cboxInstance.SpaceFind(space.Selector.NamespaceType, space.Selector.Namespace, space.Label)
	if err == nil {
		t.Error("space still found after deleting it")
	}

}

func TestSpaceLabelUniquenessOnCreation(t *testing.T) {
	cboxInstance := initializeCBox()

	s1 := createSpace(t, cboxInstance)

	s2 := models.Space{
		Label:       s1.Label,
		Description: randString(15),
	}

	// clone the same selector
	s2.Selector = s1.Selector.CloneForItem("")

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

	expected := []string{"@default", s1.String(), s3.String()}

	cboxInstance.SpaceDestroy(s2)
	core.DeleteSpaceFile(s2.Selector)

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
		t.Fatalf("space search error: %v", err)
	}

	if len(result) != 1 || result[0].Label != c1.Label {
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

	if len(result) != 1 || result[0].Label != c1.Label {
		t.Errorf("space search did not return the expected command: %v", c1)
	}

	result, err = s1.SearchCommands("non-existing", criteria)

	if err != nil {
		t.Fatalf("space search error (non-existing): %v", err)
	}

	if len(result) != 0 {
		t.Errorf("space search did return something for non-existing tag: %v", result)
	}
}
