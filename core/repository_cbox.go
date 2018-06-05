package core

import (
	"github.com/dpecos/cbox/models"
)

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
