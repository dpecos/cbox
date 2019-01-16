package tests

import (
	"testing"

	"github.com/dplabs/cbox/src/models"
)

func TestSpaceCreationNoCloudLogin(t *testing.T) {
	cboxInstance := initializeCBox()

	space := models.Space{
		// Namespace:   "",
		Label:       randString(8),
		Description: randString(15),
	}

	err := cboxInstance.SpaceCreate(&space)
	if err != nil {
		t.Fatal(err)
	}

	cboxInstance = reloadCBox(cboxInstance)

	s, err := cboxInstance.SpaceFind("", space.Label)
	if err != nil {
		t.Fatalf("could not find space: %v", err)
	}

	if s.String() != space.String() {
		t.Errorf("created space and reloaded space don't match: %s - %s", s.String(), space.String())
	}

	assertSpaceFileExists(t, cboxInstance, &space)
}

func TestSpaceCreationWithCloudLogin(t *testing.T) {
	cboxInstance := initializeCBox()

	space := createSpace(t, cboxInstance)

	cboxInstance = reloadCBox(cboxInstance)

	s, err := cboxInstance.SpaceFind(space.Namespace, space.Label)
	if err != nil {
		t.Fatalf("could not find space: %v", err)
	}

	if s.String() != space.String() {
		t.Errorf("created space and reloaded space don't match: %s - %s", s.String(), space.String())
	}

	assertSpaceFileExists(t, cboxInstance, space)
}

func TestSpaceCreationSameLabelDifferentNamespace(t *testing.T) {
	cboxInstance := initializeCBox()

	space1 := createSpace(t, cboxInstance)

	space2 := models.Space{
		Namespace: "diff-" + space1.Namespace,
		Label:     space1.Label,
	}

	err := cboxInstance.SpaceCreate(&space2)
	if err != nil {
		t.Fatal(err)
	}

	cboxInstance = reloadCBox(cboxInstance)

	s1, err := cboxInstance.SpaceFind(space1.Namespace, space1.Label)
	if err != nil {
		t.Fatalf("could not find space: %v", err)
	}

	if s1.String() != space1.String() {
		t.Errorf("created space and reloaded space don't match: %s - %s", s1.String(), space1.String())
	}

	s2, err := cboxInstance.SpaceFind(space2.Namespace, space2.Label)
	if err != nil {
		t.Fatalf("could not find space: %v", err)
	}

	if s2.String() != space2.String() {
		t.Errorf("created space and reloaded space don't match: %s - %s", s2.String(), space2.String())
	}

	assertSpaceFileExists(t, cboxInstance, space1)
	assertSpaceFileExists(t, cboxInstance, &space2)
}
