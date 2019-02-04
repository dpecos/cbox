package controllers

import (
	"fmt"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
)

func (ctrl *CLIController) findSpace(selector *models.Selector) (*models.Space, error) {

	if selector == nil {
		return nil, fmt.Errorf("find space: nil selector")
	}

	space, err := ctrl.cbox.SpaceFind(selector.NamespaceType, selector.Namespace, selector.Space)

	// if not namespace specified, maybe belongs to the logged in user
	if selector.NamespaceType != models.TypeNone || err == nil {
		return space, err
	}

	if ctrl.cloud != nil {
		return ctrl.cbox.SpaceFind(models.TypeUser, ctrl.cloud.Login, selector.Space)
	} else {
		return nil, err
	}
}

func (ctrl *CLIController) cleanOldSpaceFile(space *models.Space, oldSelector *models.Selector) {
	if space.Label != oldSelector.Space || (oldSelector.Namespace != "" && space.Selector.Namespace != oldSelector.Namespace) {
		core.DeleteSpaceFile(oldSelector)
	}
}
