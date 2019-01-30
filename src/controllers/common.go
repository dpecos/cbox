package controllers

import (
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

func InitController() *CLIController {

	core.LoadSettings("")

	cbox := core.LoadCbox()
	cloud := core.CloudClient()

	controller := CLIController{
		cbox,
		cloud,
	}

	return &controller
}
