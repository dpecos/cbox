package tests

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
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

func cloudConnect(jwt string) *core.Cloud {
	_, _, _, err := core.CloudLogin(jwt)
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}

	cloud, err := core.CloudClient()
	if err != nil {
		log.Fatalf("could create cloud client: %v", err)
	}

	return cloud
}

func reloadCBox(cbox *models.CBox) *models.CBox {
	if cbox != nil {
		core.Save(cbox)
	}

	return core.Load()
}
