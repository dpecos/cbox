package core

import (
	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/repository"
)

func LoadCbox() *models.CommandBox {

	cbox := models.CommandBox{
		Spaces: []models.Space{},
	}

	spaceIds := repository.SpacesList()

	for _, spaceId := range spaceIds {
		space := repository.SpaceLoad(spaceId)
		cbox.SpaceAdd(*space)
	}

	return &cbox
}

func PersistCbox(cbox *models.CommandBox) {
	for _, space := range cbox.Spaces {
		repository.SpaceStore(space)
	}
}
