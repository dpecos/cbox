package tests

import (
	"math/rand"
	"os"
	"time"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
)

var (
	cbox *models.CBox
)

func setupTests() {
	os.RemoveAll("/tmp/.cbox")

	reloadCBox()
}

func reloadCBox() {
	if cbox != nil {
		core.PersistCbox(cbox)
	}

	core.CheckCboxDir("/tmp")
	cbox = core.LoadCbox("/tmp")
}

func init() {
	rand.Seed(time.Now().UnixNano())
	setupTests()
}
