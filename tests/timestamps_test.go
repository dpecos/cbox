package tests

import (
	"testing"
	"time"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
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

	space := createSpace(t)
	tc := space.CreatedAt
	tu := space.UpdatedAt

	if !tc.Equal(tu) {
		t.Errorf("creation and update timestamsp should be the same for new spaces")
	}

	reloadCBox()
	s, _ := cbox.SpaceFind(space.ID.String())

	if !tc.Equal(s.CreatedAt) {
		t.Errorf("space creation timestamp changed after persisting and reloading cbox: '%s' - '%s'", tc.StringRaw(), s.CreatedAt.StringRaw())
	}

	if !tu.Equal(s.UpdatedAt) {
		t.Errorf("space update timestamp changed after persisting and reloading cbox: '%s' - '%s'", tu.StringRaw(), s.UpdatedAt.StringRaw())
	}

	s.Label = s.Label + "-updated"

	err := cbox.SpaceEdit(s, space.Label)
	if err != nil {
		t.Errorf("failed to rename space: %v", err)
	}

	spaceToDelete := &models.Space{
		Label: space.Label,
	}
	core.SpaceDeleteFile(spaceToDelete)

	reloadCBox()
	s, _ = cbox.SpaceFind(space.ID.String())

	if !tc.Equal(s.CreatedAt) {
		t.Errorf("space creation timestamp changed after update: '%s' - '%s'", tc.StringRaw(), s.CreatedAt.StringRaw())
	}

	if !s.UpdatedAt.After(tu) {
		t.Errorf("space update timestamp did not change after update: '%s' - '%s'", tu.StringRaw(), s.UpdatedAt.StringRaw())
	}
}

func TestCommandTimestamps(t *testing.T) {
	setupMocks()

	space := createSpace(t)
	tsc := space.CreatedAt
	tsu := space.UpdatedAt

	cmd := createCommand(t, space)
	tcc := cmd.CreatedAt
	tcu := cmd.UpdatedAt

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

	err := space.CommandEdit(cmd, previousLabel)
	if err != nil {
		t.Errorf("failed to rename command: %v", err)
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
