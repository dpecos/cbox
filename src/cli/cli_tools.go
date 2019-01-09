package cli

import (
	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
)

func findSpace(selector *models.Selector) (*models.Space, error) {
	namespace := selector.User
	if namespace == "" {
		namespace = cloud.Login
	}

	return cboxInstance.SpaceFind(namespace, selector.Space)
}

func cleanOldSpaceFile(space *models.Space, selector *models.Selector) {
	if space.Label != selector.Space || (selector.User != "" && space.Namespace != selector.User) {
		spaceToDelete := models.Space{
			Namespace: selector.User,
			Label:     selector.Space,
		}
		core.DeleteSpaceFile(&spaceToDelete)
	}
}
