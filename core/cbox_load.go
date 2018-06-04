package core

import (
	"github.com/dpecos/cbox/models"
	"github.com/dpecos/cbox/repository"
)

func LoadCbox() models.CommandBox {

	cbox := models.CommandBox{
		Spaces: []models.Space{},
	}

	spaceIds := repository.SpacesList()

	for _, spaceId := range spaceIds {
		space := repository.SpaceLoad(spaceId)
		cbox.SpaceAdd(space)
	}

	return cbox
}
