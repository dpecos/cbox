package cli

import (
	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
)

func findSpace(selector *models.Selector) (*models.Space, error) {
	namespace := selector.Namespace

	space, err := cboxInstance.SpaceFind(namespace, selector.Space)

	if namespace != "" || err == nil {
		return space, err
	}

	namespace = cloud.Login

	return cboxInstance.SpaceFind(namespace, selector.Space)
}

func cleanOldSpaceFile(space *models.Space, selector *models.Selector) {
	if space.Label != selector.Space || (selector.Namespace != "" && space.Namespace != selector.Namespace) {
		spaceToDelete := models.Space{
			Namespace: selector.Namespace,
			Label:     selector.Space,
		}
		core.DeleteSpaceFile(&spaceToDelete)
	}
}
