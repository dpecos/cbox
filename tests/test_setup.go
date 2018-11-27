package tests

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/dplabs/cbox/internal/app/core"
	"github.com/dplabs/cbox/pkg/models"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func initializeCBox() *models.CBox {
	if err := os.RemoveAll("/tmp/.cbox"); err != nil {
		log.Fatalf("could not clean cbox test directory: %v", err)
	}

	core.LoadSettings("/tmp")
	return reloadCBox(nil)
}

func reloadCBox(cbox *models.CBox) *models.CBox {
	if cbox != nil {
		core.Save(cbox)
	}

	return core.Load()
}
