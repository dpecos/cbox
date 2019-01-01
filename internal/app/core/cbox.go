package core

import (
	"log"

	"github.com/dplabs/cbox/pkg/models"
)

const (
	defaultSpaceID          = "default"
	defaultSpaceDescription = "Default space to store commands"

	pathSpaces     = "spaces"
	pathConfigFile = "config.yml"
)

func Load() *models.CBox {
	cbox := models.CBox{
		Spaces: []*models.Space{},
	}

	spaces := spacesLoad()

	for _, space := range spaces {
		err := cbox.SpaceCreate(space)
		if err != nil {
			log.Fatalf("load: could not create space: %v", err)
		}
	}
	return &cbox
}

func Save(cbox *models.CBox) {
	for _, space := range cbox.Spaces {
		spaceStoreFile(space)
	}
}

func DeleteSpaceFile(space *models.Space) {
	spaceDeleteFile(space)
}
