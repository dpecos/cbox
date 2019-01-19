package cli

import (
	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
)

func findSpace(selector *models.Selector) (*models.Space, error) {

	space, err := cboxInstance.SpaceFind(selector.NamespaceType, selector.Namespace, selector.Space)

	// if not namespace specified, maybe belongs to the logged in user
	if selector.NamespaceType != models.TypeNone || err == nil {
		return space, err
	}

	return cboxInstance.SpaceFind(models.TypeUser, cloud.Login, selector.Space)
}

func cleanOldSpaceFile(space *models.Space, oldSelector *models.Selector) {
	if space.Label != oldSelector.Space || (oldSelector.Namespace != "" && space.Selector.Namespace != oldSelector.Namespace) {
		core.DeleteSpaceFile(oldSelector)
	}
}
