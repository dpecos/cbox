package controllers

import (
	"fmt"
	"log"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
)

func (ctrl *CLIController) parseSelectorAllowEmpty(args []string) *models.Selector {
	var selectorStr = ""
	if len(args) == 1 {
		selectorStr = args[0]
	}

	selector, err := models.ParseSelector(selectorStr)
	if err != nil {
		log.Fatalf("parse selector empty: %v", err)
	}

	return selector
}

func (ctrl *CLIController) parseSelector(args []string) *models.Selector {
	selector, err := models.ParseSelector(args[0])
	if err != nil {
		log.Fatalf("parse selector: %v", err)
	}
	return selector
}

func (ctrl *CLIController) findSpace(selector *models.Selector) (*models.Space, error) {

	if selector == nil {
		return nil, fmt.Errorf("find space: nil selector")
	}

	space, err := ctrl.cbox.SpaceFind(selector.NamespaceType, selector.Namespace, selector.Space)

	// if not namespace specified, maybe belongs to the logged in user
	if selector.NamespaceType != models.TypeNone || err == nil {
		return space, err
	}

	return ctrl.cbox.SpaceFind(models.TypeUser, ctrl.cloud.Login, selector.Space)
}

func (ctrl *CLIController) cleanOldSpaceFile(space *models.Space, oldSelector *models.Selector) {
	if space.Label != oldSelector.Space || (oldSelector.Namespace != "" && space.Selector.Namespace != oldSelector.Namespace) {
		core.DeleteSpaceFile(oldSelector)
	}
}
