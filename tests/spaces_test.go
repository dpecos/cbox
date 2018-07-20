package tests

import (
	"os"
	"testing"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
	"github.com/satori/go.uuid"
)

func TestMain(m *testing.M) {
	core.CheckCboxDir("/tmp")
	cbox = core.LoadCbox("/tmp")

	res := m.Run()

	os.RemoveAll("/tmp/.cbox")

	os.Exit(res)
}
func TestSpaceCreationDeletion(t *testing.T) {
	id, _ := uuid.NewV4()
	label := randString(8)

	space := models.Space{
		ID:          id,
		Label:       label,
		Description: randString(15),
	}

	err := cbox.SpaceAdd(&space)

	if err != nil {
		t.Error(err)
	}

	core.PersistCbox(cbox)

	assertSpaceFileExists(t, &space)

	core.SpaceDelete(&space)

	assertSpaceFileNotExists(t, &space)
}
