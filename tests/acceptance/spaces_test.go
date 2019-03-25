package acceptance_tests

import (
	"os"
	"testing"

	"github.com/dplabs/cbox/src/tools/tty"
	"github.com/dplabs/cbox/tests"
)

func TestDefaultSpaceCreatedOnNewSetup(t *testing.T) {
	ctrl, dir := tests.InitController()
	defer os.RemoveAll(dir)

	tty.MockedOutput = ""
	ctrl.SpacesList()

	tests.AssertOutputContains(t, "@default - Default space to store commands", "default space not found")

	ctrl.ConfigGet("cbox.environment")
	tests.AssertOutputContains(t, "cbox.environment -> test", "environment setting not found")

	ctrl.ConfigGet("cbox.default-space")
	tests.AssertOutputContains(t, "cbox.default-space -> default", "default space setting not found")
}

func TestSpaceListingWhenNoSpace(t *testing.T) {
	ctrl, dir := tests.InitController()
	defer os.RemoveAll(dir)

	tty.MockedOutput = ""
	ctrl.SpacesDestroy("@default")

	tty.MockedOutput = ""
	ctrl.SpacesList()

	if tty.MockedOutput != "" {
		t.Errorf("there was output while listing empty space list: %s", tty.MockedOutput)
	}
}

func TestCreateEditDestroySpaces(t *testing.T) {
	ctrl, dir := tests.InitController()
	defer os.RemoveAll(dir)

	tty.MockedOutput = ""
	ctrl.SpacesDestroy("@default")

	tty.MockedOutput = ""
	tty.MockedInput = []string{"test-space", "This is a test space"}
	ctrl.SpacesCreate()
	tests.AssertOutputContains(t, "@test-space - This is a test space", "space creation failed")

	tty.MockedOutput = ""
	tty.MockedInput = []string{"test-space-renamed", "This is a renamed test space", "y"}
	ctrl.SpacesEdit("@test-space")
	tests.AssertOutputContains(t, "@test-space-renamed - This is a renamed test space", "space edition failed")

	tty.MockedOutput = ""
	ctrl.SpacesDestroy("@test-space-renamed")
	tests.AssertOutputContains(t, "Space destroyed successfully", "space deletion failed")

	tty.MockedOutput = ""
	ctrl.SpacesList()
	if tty.MockedOutput != "" {
		t.Errorf("space deletion left some spaces behind: %s", tty.MockedOutput)
	}
}
