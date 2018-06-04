package models

import (
	"testing"
)

func expectSelector(t *testing.T, s *Selector, err error, id string, tag string, space string) {
	if err != nil {
		t.Error("Error parsing selector", err)
	}
	if s.ID != id {
		t.Errorf("Expecting ID '%s' but got '%s'", id, s.ID)
	}
	if s.Tag != tag {
		t.Errorf("Expecting Tag '%s' but got '%s'", tag, s.Tag)
	}
	if s.Space != space {
		t.Errorf("Expecting Space '%s' but got '%s'", space, s.Space)
	}
}

func TestEmptySelector(t *testing.T) {
	s, err := ParseSelector("")
	expectSelector(t, s, err, "", "", "default")
}

func TestSpaceSelector(t *testing.T) {
	s, err := ParseSelector("@test")
	expectSelector(t, s, err, "", "", "test")
}
func TestTagSelector(t *testing.T) {
	s, err := ParseSelector("#test-tag")
	expectSelector(t, s, err, "", "test-tag", "default")
}

func TestTagSpaceSelector(t *testing.T) {
	s, err := ParseSelector("#test-tag@test")
	expectSelector(t, s, err, "", "test-tag", "test")
}

func TestIDSelector(t *testing.T) {
	s, err := ParseSelector("test-id")
	expectSelector(t, s, err, "test-id", "", "default")
}

func TestIDSpaceSelector(t *testing.T) {
	s, err := ParseSelector("test-id@test")
	expectSelector(t, s, err, "test-id", "", "test")
}

func TestInvalidIdTagSelector(t *testing.T) {
	_, err := ParseSelector("test-id#test-tag")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestInvalidTwoTagSelector(t *testing.T) {
	_, err := ParseSelector("#test-tag#test-tag")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestInvalidEmptyTagSelector(t *testing.T) {
	_, err := ParseSelector("#")

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
