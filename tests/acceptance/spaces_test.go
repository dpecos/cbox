package acceptance_tests

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/dplabs/cbox/src/controllers"
	"github.com/dplabs/cbox/src/tools/tty"
)

func TestDefaultSpaceCreatedOnNewSetup(t *testing.T) {
	dir, _ := ioutil.TempDir("", "cbox")
	defer os.RemoveAll(dir)
	ctrl := controllers.InitController(dir)

	tty.MockedOutput = ""
	ctrl.SpacesList()

	checkOutput(t, "@default - Default space to store commands", "default space not found")

	ctrl.ConfigGet("cbox.environment")
	checkOutput(t, "cbox.environment -> test", "environment setting not found")

	ctrl.ConfigGet("cbox.default-space")
	checkOutput(t, "cbox.default-space -> default", "default space setting not found")
}

func TestSpaceListingWhenNoSpace(t *testing.T) {
	dir, _ := ioutil.TempDir("", "cbox")
	defer os.RemoveAll(dir)
	ctrl := controllers.InitController(dir)

	tty.MockedOutput = ""
	ctrl.SpacesDestroy("@default")

	tty.MockedOutput = ""
	ctrl.SpacesList()

	if tty.MockedOutput != "" {
		t.Errorf("there was output while listing empty space list: %s", tty.MockedOutput)
	}
}

func TestCreateEditDestroySpaces(t *testing.T) {
	dir, _ := ioutil.TempDir("", "cbox")
	defer os.RemoveAll(dir)
	ctrl := controllers.InitController(dir)

	tty.MockedOutput = ""
	ctrl.SpacesDestroy("@default")

	tty.MockedOutput = ""
	tty.MockedInput = []string{"test-space", "This is a test space"}
	ctrl.SpacesCreate()
	checkOutput(t, "@test-space - This is a test space", "space creation failed")

	tty.MockedOutput = ""
	tty.MockedInput = []string{"test-space-renamed", "This is a renamed test space", "y"}
	ctrl.SpacesEdit("@test-space")
	checkOutput(t, "@test-space-renamed - This is a renamed test space", "space edition failed")

	tty.MockedOutput = ""
	ctrl.SpacesDestroy("@test-space-renamed")
	checkOutput(t, "Space destroyed successfully", "space deletion failed")

	tty.MockedOutput = ""
	ctrl.SpacesList()
	if tty.MockedOutput != "" {
		t.Errorf("space deletion left some spaces behind: %s", tty.MockedOutput)
	}
}
