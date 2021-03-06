package models_test

import (
	"os"
	"testing"

	"github.com/dplabs/cbox/src/models"
	"github.com/spf13/viper"
)

func TestMain(m *testing.M) {
	viper.SetDefault("cbox.default-space", "default")

	os.Exit(m.Run())
}

func expectSelector(t *testing.T, s *models.Selector, err error, item string, org string, user string, space string) {
	if err != nil {
		t.Fatal("Error parsing selector", err)
	}
	if s == nil {
		t.Fatal("Nil selector passed")
	}
	if s.Item != item {
		t.Errorf("Expecting ID '%s' but got '%s'", item, s.Item)
	}
	if s.Space != space {
		t.Errorf("Expecting Space '%s' but got '%s'", space, s.Space)
	}
}

func TestEmptySelector(t *testing.T) {
	s, err := models.ParseSelector("")
	expectSelector(t, s, err, "", "", "", "default")
}

func TestSpaceSelector(t *testing.T) {
	s, err := models.ParseSelector("@test")
	expectSelector(t, s, err, "", "", "", "test")
}
func TestTagSelector(t *testing.T) {
	s, err := models.ParseSelector("test-tag")
	expectSelector(t, s, err, "test-tag", "", "", "default")
}

func TestTagSpaceSelector(t *testing.T) {
	s, err := models.ParseSelector("test-tag@test")
	expectSelector(t, s, err, "test-tag", "", "", "test")
}

func TestMandatorySpaceSelector(t *testing.T) {
	s, err := models.ParseSelectorMandatorySpace("@test")
	expectSelector(t, s, err, "", "", "", "test")
}

func TestMandatoryItemSelector(t *testing.T) {
	s, err := models.ParseSelectorMandatoryItem("item@test")
	expectSelector(t, s, err, "item", "", "", "test")
}

func TestUserSpaceSelector(t *testing.T) {
	s, err := models.ParseSelectorMandatoryItem("item@user:space")
	expectSelector(t, s, err, "item", "", "user", "space")
}

func TestOrgSpaceSelector(t *testing.T) {
	s, err := models.ParseSelectorMandatoryItem("item@org/space")
	expectSelector(t, s, err, "item", "org", "", "space")
}

func TestInvalidOrgUserSpaceSelector(t *testing.T) {
	_, err := models.ParseSelectorMandatoryItem("item@org/user:space")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestEmptyMandatorySpaceSelector(t *testing.T) {
	_, err := models.ParseSelectorMandatorySpace("")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestEmptyMandatoryItem(t *testing.T) {
	_, err := models.ParseSelectorMandatoryItem("@space")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestInvalidIdTagSelector(t *testing.T) {
	_, err := models.ParseSelector("invalid%chars")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestInvalidEmptySpaceSelector(t *testing.T) {
	_, err := models.ParseSelector("@")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestInvalidUppercaseSpaceSelector(t *testing.T) {
	_, err := models.ParseSelector("Test")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestInvalidCharacterInIDSelector(t *testing.T) {
	_, err := models.ParseSelector("t/@space")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestInvalidCharacter1InSpaceSelector(t *testing.T) {
	_, err := models.ParseSelector("t@space/d/s")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestInvalidCharacter2InSpaceSelector(t *testing.T) {
	_, err := models.ParseSelector("t@space:d/s")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestInvalidCharacter3InSpaceSelector(t *testing.T) {
	_, err := models.ParseSelector("t@space:d:s")

	if err == nil {
		t.Error("Expected error was not created")
	}
}

func TestStringWithUser(t *testing.T) {
	sel := "item@user:space"
	s, _ := models.ParseSelectorMandatoryItem(sel)
	str := s.String()

	if sel != str {
		t.Errorf("Generated string does not match expected value: expected = '%s', got = '%s'", sel, str)
	}
}

func TestStringWithoutUser(t *testing.T) {
	sel := "item@space"
	s, _ := models.ParseSelectorMandatoryItem(sel)
	str := s.String()

	if sel != str {
		t.Errorf("Generated string does not match expected value: expected = '%s', got = '%s'", sel, str)
	}
}

func TestStringWithoutItem(t *testing.T) {
	sel := "@space"
	s, _ := models.ParseSelectorMandatorySpace(sel)
	str := s.String()

	if sel != str {
		t.Errorf("Generated string does not match expected value: expected = '%s', got = '%s'", sel, str)
	}
}
