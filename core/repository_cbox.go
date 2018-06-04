package core

import (
	"github.com/dpecos/cbox/models"
)

func LoadCbox() *models.CommandBox {

	cbox := models.CommandBox{
		Spaces: []models.Space{},
	}

	spaceIds := SpacesList()

	for _, spaceId := range spaceIds {
		space := SpaceLoad(spaceId)
		cbox.SpaceAdd(*space)
	}

	return &cbox
}

func PersistCbox(cbox *models.CommandBox) {
	for _, space := range cbox.Spaces {
		SpaceStore(space)
	}
}
