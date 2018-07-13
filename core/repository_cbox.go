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
		log.Fatalf("respository: could not get HOME: %v", err)
	}
	return path.Join(home, CBOX_PATH, content)
}

func CheckCboxDir() string {
	cboxPath := resolveInCboxDir("")
	tools.CreateDirectoryIfNotExists(cboxPath)

	configFile := resolveInCboxDir("config.yml")
	tools.CreateFileIfNotExists(configFile)

	spacesPath := resolveInCboxDir("spaces")
	if tools.CreateDirectoryIfNotExists(spacesPath) {
		defaultSpace := models.Space{
			ID:          DEFAULT_SPACE_ID,
			Description: DEFAULT_SPACE_DESCRIPTION,
		}
		cbox := LoadCbox()
		cbox.SpaceCreate(&defaultSpace)
		PersistCbox(cbox)
	}

	return cboxPath
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
