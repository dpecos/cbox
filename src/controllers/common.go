package controllers

import (
	"log"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
)

var (
	SkipQuestionsFlag      bool
	ShowCommandsSourceFlag bool
	SourceOnlyFlag         bool
	OrganizationOption     string
)

type CLIController struct {
	cbox  *models.CBox
	cloud *core.Cloud
}

func InitController(cbox *models.CBox, cloud *core.Cloud) *CLIController {
	if cbox == nil {
		log.Fatalf("cbox instance not loaded")
	}
	if cloud == nil {
		log.Fatalf("cloud config not loaded")
	}

	controller := CLIController{
		cbox,
		cloud,
	}

	return &controller
}
