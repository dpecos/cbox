package models

import (
	"testing"

	"github.com/gofrs/uuid"
)

func TestCommandMatches(t *testing.T) {
	id, _ := uuid.NewV4()
	command := Command{
		Label:       "label",
		Description: "description",
		Code:        "code",
	}
	command.ID = id

	if command.Matches(id.String()[0:5]) {
		t.Errorf("command match by ID and shouldn't")
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
