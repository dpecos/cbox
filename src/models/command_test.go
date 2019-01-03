package models

import (
	"testing"
)

func TestCommandMatches(t *testing.T) {
	command := Command{
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
