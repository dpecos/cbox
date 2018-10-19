package models

import (
	"testing"

	"github.com/spf13/viper"
)

func TestMain(m *testing.M) {
	viper.SetDefault("cbox.default-space", "default")
}

func expectSelector(t *testing.T, s *Selector, err error, item string, org string, user string, space string) {
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
	expectSelector(t, s, err, "", "", "", "default")
}

func TestSpaceSelector(t *testing.T) {
	s, err := ParseSelector("@test")
	expectSelector(t, s, err, "", "", "", "test")
}
func TestTagSelector(t *testing.T) {
	s, err := ParseSelector("test-tag")
	expectSelector(t, s, err, "test-tag", "", "", "default")
}

func TestTagSpaceSelector(t *testing.T) {
	s, err := ParseSelector("test-tag@test")
	expectSelector(t, s, err, "test-tag", "", "", "test")
}

func TestMandatorySpaceSelector(t *testing.T) {
	s, err := ParseSelectorMandatorySpace("@test")
	expectSelector(t, s, err, "", "", "", "test")
}

func TestMandatoryItemSelector(t *testing.T) {
	s, err := ParseSelectorMandatoryItem("item@test")
	expectSelector(t, s, err, "item", "", "", "test")
}

func TestUserSpaceSelector(t *testing.T) {
	s, err := ParseSelectorMandatoryItem("item@user:space")
	expectSelector(t, s, err, "item", "", "user", "space")
}

func TestOrgSpaceSelector(t *testing.T) {
	s, err := ParseSelectorMandatoryItem("item@org/space")
	expectSelector(t, s, err, "item", "org", "", "space")
}

func TestOrgUserSpaceSelector(t *testing.T) {
	s, err := ParseSelectorMandatoryItem("item@org/user:space")
	expectSelector(t, s, err, "item", "org", "user", "space")
}

func TestEmptyMandatorySpaceSelector(t *testing.T) {
	_, err := ParseSelectorMandatorySpace("")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestEmptyMandatoryItem(t *testing.T) {
	_, err := ParseSelectorMandatoryItem("@space")

	if err == nil {
		t.Error("Expected error was not created")
	}
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

func TestInvalidCharacterInIDSelector(t *testing.T) {
	_, err := ParseSelector("t/@space")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestInvalidCharacter1InSpaceSelector(t *testing.T) {
	_, err := ParseSelector("t@space/d/s")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestInvalidCharacter2InSpaceSelector(t *testing.T) {
	_, err := ParseSelector("t@space:d/s")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestInvalidCharacter3InSpaceSelector(t *testing.T) {
	_, err := ParseSelector("t@space:d:s")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestStringWithUser(t *testing.T) {
	sel := "item@user:space"
	s, _ := ParseSelectorMandatoryItem(sel)
	str := s.String()

	if sel != str {
		t.Errorf("Generated string does not match expected value: expected = '%s', got = '%s'", sel, str)
	}
}

func TestStringWithoutUser(t *testing.T) {
	sel := "item@space"
	s, _ := ParseSelectorMandatoryItem(sel)
	str := s.String()

	if sel != str {
		t.Errorf("Generated string does not match expected value: expected = '%s', got = '%s'", sel, str)
	}
}

func TestStringWithoutItem(t *testing.T) {
	sel := "@space"
	s, _ := ParseSelectorMandatorySpace(sel)
	str := s.String()

	if sel != str {
		t.Errorf("Generated string does not match expected value: expected = '%s', got = '%s'", sel, str)
	}
}
