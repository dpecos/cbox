package tests

import (
	"testing"

	"github.com/dpecos/cbox/core"
)

func TestSpaceCreationDeletion(t *testing.T) {

	space = createSpace(t)

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

	core.SpaceDelete(space)

	assertSpaceFileNotExists(t, space)
}
