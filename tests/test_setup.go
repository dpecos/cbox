package tests

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/dpecos/cbox/internal/app/core"
	"github.com/dpecos/cbox/pkg/models"
	"github.com/spf13/viper"
)

var (
	cbox *models.CBox
)

func init() {
	rand.Seed(time.Now().UnixNano())
	setupTests()
}

func setupTests() {
	cbox = nil

	os.RemoveAll("/tmp/.cbox")

	core.InitCBox("/tmp")
	reloadCBox()

	fmt.Println(viper.GetString("config.default-space"))
}

func reloadCBox() {
	if cbox != nil {
		core.PersistCbox(cbox)
	}

	core.CheckCboxDir("/tmp")
	cbox = core.LoadCbox("/tmp")
}
