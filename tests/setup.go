package tests

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/dplabs/cbox/src/controllers"
	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools/tty"
)

func setupTestEnv() {
	os.Setenv("CBOX_ENV", "test") // TODO check if needed

	tty.DisableColors = true
	// tty.DisableOutput = true
	tty.MockTTY = true
	tty.SkipQuestions = true

	rand.Seed(time.Now().UnixNano())
}

func InitController() (*controllers.CLIController, string) {
	setupTestEnv()

	dir, _ := ioutil.TempDir("", "cbox")
	ctrl := controllers.InitController(dir)

	ctrl.ConfigSet("cbox.environment", "test") // TODO check if needed

	tty.MockedInput = []string{}
	tty.MockedOutput = ""

	return ctrl, dir
}

func InitializeCBox() *models.CBox {
	setupTestEnv()

	if err := os.RemoveAll("/tmp/.cbox"); err != nil {
		log.Fatalf("could not clean cbox test directory: %v", err)
	}

	return core.Load("/tmp")
}

func ReloadCBox(cbox *models.CBox) *models.CBox {
	if cbox != nil {
		core.Save(cbox)
	}

	return core.Load("/tmp")
}
