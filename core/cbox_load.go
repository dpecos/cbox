package core

import (
	"github.com/dpecos/cmdbox/models"
	"github.com/dpecos/cmdbox/repository"
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
