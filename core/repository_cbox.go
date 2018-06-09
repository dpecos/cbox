package core

import (
	"log"
	"path"

	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/tools"
	homedir "github.com/mitchellh/go-homedir"
)

const CBOX_PATH = ".cbox"

func resolveInCboxDir(content string) string {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Could not retrieve HOME: %s", err)
	}
	return path.Join(home, CBOX_PATH, content)
}

func CheckCboxDir() {
	cboxPath := resolveInCboxDir("")
	tools.CreateDirectoryIfNotExists(cboxPath)

	spacesPath := resolveInCboxDir("spaces")
	if tools.CreateDirectoryIfNotExists(spacesPath) {
		defaultSpace := models.Space{
			Name:  DEFAULT_SPACE_NAME,
			Title: DEFAULT_SPACE_TITLE,
		}
		cbox := LoadCbox()
		cbox.SpaceCreate(&defaultSpace)
		PersistCbox(cbox)
	}
}

func LoadCbox() *models.CBox {

	cbox := models.CBox{
		Spaces: []models.Space{},
	}

	spaces := SpaceList()

	for _, space := range spaces {
		cbox.SpaceAdd(space)
	}

	return &cbox
}

func PersistCbox(cbox *models.CBox) {
	for _, space := range cbox.Spaces {
		SpaceStore(&space)
	}
}
