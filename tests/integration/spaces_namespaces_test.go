package integration_tests

import (
	"testing"

	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/tests"
)

func TestSpaceCreationNoCloudLogin(t *testing.T) {
	cboxInstance := tests.InitializeCBox()

	space := models.Space{
		Label:       tests.RandString(8),
		Description: tests.RandString(15),
	}
	space.Selector = models.NewSelector(models.TypeNone, "", space.Label, "")

	err := cboxInstance.SpaceCreate(&space)
	if err != nil {
		t.Fatal(err)
	}

	cboxInstance = tests.ReloadCBox(cboxInstance)

	s, err := cboxInstance.SpaceFind(models.TypeNone, "", space.Label)
	if err != nil {
		t.Fatalf("could not find space: %v", err)
	}

	if s.String() != space.String() {
		t.Errorf("created space and reloaded space don't match: %s - %s", s.String(), space.String())
	}

	tests.AssertSpaceFileExists(t, cboxInstance, &space)
}

func TestSpaceCreationWithCloudLogin(t *testing.T) {
	cboxInstance := tests.InitializeCBox()

	space := tests.CreateSpace(t, cboxInstance)

	cboxInstance = tests.ReloadCBox(cboxInstance)

	s, err := cboxInstance.SpaceFind(space.Selector.NamespaceType, space.Selector.Namespace, space.Label)
	if err != nil {
		t.Fatalf("could not find space: %v", err)
	}

	if s.String() != space.String() {
		t.Errorf("created space and reloaded space don't match: %s - %s", s.String(), space.String())
	}

	tests.AssertSpaceFileExists(t, cboxInstance, space)
}

func TestSpaceCreationSameLabelDifferentNamespace(t *testing.T) {
	cboxInstance := tests.InitializeCBox()

	space1 := tests.CreateSpace(t, cboxInstance)

	space2 := models.Space{
		Label: space1.Label,
	}
	space2.Selector = models.NewSelector(space1.Selector.NamespaceType, "diff-"+space1.Selector.Namespace, space2.Label, "")

	err := cboxInstance.SpaceCreate(&space2)
	if err != nil {
		t.Fatal(err)
	}

	cboxInstance = tests.ReloadCBox(cboxInstance)

	s1, err := cboxInstance.SpaceFind(space1.Selector.NamespaceType, space1.Selector.Namespace, space1.Label)
	if err != nil {
		t.Fatalf("could not find space: %v", err)
	}

	if s1.String() != space1.String() {
		t.Errorf("created space and reloaded space don't match: %s - %s", s1.String(), space1.String())
	}

	s2, err := cboxInstance.SpaceFind(space2.Selector.NamespaceType, space2.Selector.Namespace, space2.Label)
	if err != nil {
		t.Fatalf("could not find space: %v", err)
	}

	if s2.String() != space2.String() {
		t.Errorf("created space and reloaded space don't match: %s - %s", s2.String(), space2.String())
	}

	tests.AssertSpaceFileExists(t, cboxInstance, space1)
	tests.AssertSpaceFileExists(t, cboxInstance, &space2)
}
