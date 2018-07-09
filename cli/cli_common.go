package cli

import (
	"log"

	"github.com/dpecos/cbox/models"
)

type CLIController struct {
}

var ctrl = CLIController{}

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
