package models_test

import (
	"testing"

	"github.com/dplabs/cbox/src/models"
)

func TestCommandMatches(t *testing.T) {
	command := models.Command{
		Label:       "label",
		Description: "description",
		Code:        "code",
	}

	if !command.Matches("lab") {
		t.Errorf("command did not match by Label")
	}
	if !command.Matches("desc") {
		t.Errorf("command did not match by Description")
	}
	if !command.Matches("cod") {
		t.Errorf("command did not match by Code")
	}
}
