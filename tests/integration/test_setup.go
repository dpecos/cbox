package integration_tests

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools/tty"
)

func init() {
	os.Setenv("CBOX_ENV", "test")
	tty.DisableOutput = true

	rand.Seed(time.Now().UnixNano())
}

func initializeCBox() *models.CBox {
	if err := os.RemoveAll("/tmp/.cbox"); err != nil {
		log.Fatalf("could not clean cbox test directory: %v", err)
	}

	return core.Load("/tmp")
}

func reloadCBox(cbox *models.CBox) *models.CBox {
	if cbox != nil {
		core.Save(cbox)
	}

	return core.Load("/tmp")
}
