package models

import (
	"testing"
)

func expectSelector(t *testing.T, s *Selector, err error, item string, space string) {
	if err != nil {
		t.Error("Error parsing selector", err)
	}
	if s.Item != item {
		t.Errorf("Expecting ID '%s' but got '%s'", item, s.Item)
	}
	if s.Space != space {
		t.Errorf("Expecting Space '%s' but got '%s'", space, s.Space)
	}
}

func TestEmptySelector(t *testing.T) {
	s, err := ParseSelector("")
	expectSelector(t, s, err, "", "default")
}

func TestSpaceSelector(t *testing.T) {
	s, err := ParseSelector("@test")
	expectSelector(t, s, err, "", "test")
}
func TestTagSelector(t *testing.T) {
	s, err := ParseSelector("test-tag")
	expectSelector(t, s, err, "test-tag", "default")
}

func TestTagSpaceSelector(t *testing.T) {
	s, err := ParseSelector("test-tag@test")
	expectSelector(t, s, err, "test-tag", "test")
}

func TestInvalidIdTagSelector(t *testing.T) {
	_, err := ParseSelector("invalid%chars")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestInvalidEmptySpaceSelector(t *testing.T) {
	_, err := ParseSelector("@")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestInvalidUppercaseSpaceSelector(t *testing.T) {
	_, err := ParseSelector("Test")

	if err == nil {
		t.Error("Expected error was not created")
	}
}
