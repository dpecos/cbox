package controllers

import (
	"fmt"
	"log"
	"strings"

	"github.com/dplabs/cbox/src/models"
	"github.com/dplabs/cbox/src/tools/console"
)

func (ctrl *CLIController) SearchCommands(spcSelectorStr *string, criteria string) {

	if criteria == "" {
		log.Fatal("search: criteria not specified")
	}

	if strings.Contains(criteria, "@") {
		log.Fatalf("search: criteria not specified - this looks like a selector")
	}

	sel := ""
	if spcSelectorStr != nil {
		sel = *spcSelectorStr
	}

	selector, err := models.ParseSelector(sel)
	if err != nil {
		log.Fatalf("search: %v", err)
	}

	var spaces []*models.Space = []*models.Space{}
	if sel != "" {
		space, err := ctrl.findSpace(selector)
		if err != nil {
			log.Fatalf("search: %v", err)
		}
		spaces = append(spaces, space)
	} else {
		spaces = ctrl.cbox.Spaces
	}

	var commands []*models.Command = []*models.Command{}
	for _, space := range spaces {
		var err error
		cs, err := space.SearchCommands(selector.Item, criteria)
		if err != nil {
			log.Fatalf("search: %v", err)
		}
		commands = append(commands, cs...)
	}

	header := fmt.Sprintf("Results for \"%s\"", criteria)
	if spcSelectorStr != nil {
		header = fmt.Sprintf("%s in '%s'", header, selector.String())
	}
	console.PrintCommandList(header, commands, ListingsModeOption, ListingsSortOption)
}
