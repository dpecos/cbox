package controllers

import (
	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
)

var (
	SkipQuestionsFlag      bool
	ShowCommandsSourceFlag bool
	SourceOnlyFlag         bool
	ListingsModeOption     string
	ListingsSortOption     string
	OrganizationOption     string
)

type CLIController struct {
	cbox  *models.CBox
	cloud *models.Cloud
}

func InitController(path string) *CLIController {
	cbox := core.Load(path)
	cloud := core.CloudClient(cbox)

	controller := CLIController{
		cbox,
		cloud,
	}

	return &controller
}
