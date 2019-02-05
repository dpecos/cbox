package integration_tests

import (
	"testing"
	"time"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
)

func setupMocks() {
	var timeCounter = int64(1)
	models.UnixTimeNow = func() models.UnixTime {
		t := timeCounter
		timeCounter += 1
		return models.UnixTime(time.Unix(t, 0))
	}
}

func TestSpaceTimestamps(t *testing.T) {
	setupMocks()

	cboxInstance := initializeCBox()

	space := createSpace(t, cboxInstance)
	tc := space.CreatedAt
	tu := space.UpdatedAt

	if !tc.Equal(tu) {
		t.Fatalf("creation and update timestamsp should be the same for new spaces")
	}

	cboxInstance = reloadCBox(cboxInstance)
	s, _ := cboxInstance.SpaceFind(space.Selector.NamespaceType, space.Selector.Namespace, space.Label)

	if !tc.Equal(s.CreatedAt) {
		t.Errorf("space creation timestamp changed after persisting and reloading cboxInstance: '%s' - '%s'", tc.StringRaw(), s.CreatedAt.StringRaw())
	}

	if !tu.Equal(s.UpdatedAt) {
		t.Errorf("space update timestamp changed after persisting and reloading cboxInstance: '%s' - '%s'", tu.StringRaw(), s.UpdatedAt.StringRaw())
	}

	s.Label = s.Label + "-updated"

	err := cboxInstance.SpaceEdit(s, space.Selector.Namespace, space.Label)
	if err != nil {
		t.Fatalf("failed to rename space: %v", err)
	}

	core.DeleteSpaceFile(space.Selector)

	cboxInstance = reloadCBox(cboxInstance)
	s, err = cboxInstance.SpaceFind(s.Selector.NamespaceType, s.Selector.Namespace, s.Label)

	if err != nil {
		t.Fatalf("could not find space: %v", err)
	}

	if !tc.Equal(s.CreatedAt) {
		t.Errorf("space creation timestamp changed after update: '%s' - '%s'", tc.StringRaw(), s.CreatedAt.StringRaw())
	}

	if !s.UpdatedAt.After(tu) {
		t.Errorf("space update timestamp did not change after update: '%s' - '%s'", tu.StringRaw(), s.UpdatedAt.StringRaw())
	}
}

func TestCommandTimestamps(t *testing.T) {
	setupMocks()

	cboxInstance := initializeCBox()

	space := createSpace(t, cboxInstance)
	tsc := space.CreatedAt
	tsu := space.UpdatedAt

	cmd := createCommand(t, space)
	tcc := cmd.CreatedAt
	tcu := cmd.UpdatedAt

	c, err := space.CommandFind(cmd.Label)
	if err != nil {
		t.Fatalf("could not find space after creation: %v", err)
	}

	if cmd != c {
		t.Fatalf("found space and created space are different references")
	}

	if !tsc.Equal(space.CreatedAt) {
		t.Errorf("space creation timestamp changed after new command: '%s' - '%s'", tsc.StringRaw(), space.CreatedAt.StringRaw())
	}

	if !space.UpdatedAt.After(tsu) {
		t.Errorf("space update timestamp did not change after new command: '%s' - '%s'", tsu.StringRaw(), space.UpdatedAt.StringRaw())
	}

	if !tcc.Equal(tcu) {
		t.Errorf("command creation and update timestamsp should be the same for new commands")
	}

	if !space.UpdatedAt.Equal(tcc) {
		t.Errorf("space update and last command creation timestamps should be the same")
	}

	previousLabel := cmd.Label
	cmd.Label = cmd.Label + "-update"

	err = space.CommandEdit(cmd, previousLabel)

	if err != nil {
		t.Fatalf("failed to rename command: %v", err)
	}

	if !tcc.Equal(cmd.CreatedAt) {
		t.Errorf("command creation should not change after edition : '%s' - '%s'", tcc.StringRaw(), cmd.CreatedAt.StringRaw())
	}

	if !cmd.UpdatedAt.After(tcu) {
		t.Errorf("command update should change after edition : '%s' - '%s'", tcu.StringRaw(), cmd.UpdatedAt.StringRaw())
	}

	if !cmd.UpdatedAt.Equal(space.UpdatedAt) {
		t.Errorf("space and command should have same update time after changing the command: '%s' - '%s'", cmd.UpdatedAt.StringRaw(), space.UpdatedAt.StringRaw())
	}

	tsu = space.UpdatedAt
	space.CommandDelete(cmd)

	if !space.UpdatedAt.After(tsu) {
		t.Errorf("space update time should change after deleting command: '%s' - '%s'", tsu.StringRaw(), space.UpdatedAt.StringRaw())
	}
}
