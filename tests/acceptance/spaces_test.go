package acceptance_tests

import (
	"os"
	"strings"
	"testing"

	"github.com/dplabs/cbox/src/controllers"
	"github.com/dplabs/cbox/src/tools/tty"
)

func init() {
	controllers.SkipQuestionsFlag = true
	tty.DisableColors = true
	tty.MockTTY = true
}

func checkOutput(t *testing.T, expected string, msg string) {
	if !strings.Contains(tty.MockedOutput, expected) {
		t.Errorf("%s: %s", msg, tty.MockedOutput)
	}
}

func TestDefaultSpaceCreatedOnNewSetup(t *testing.T) {
	tty.MockedOutput = ""

	os.RemoveAll("/tmp/.cbox")

	ctrl := controllers.InitController("/tmp")
	ctrl.SpacesList()

	checkOutput(t, "@default - Default space to store commands", "default space not found")

	ctrl.ConfigGet("cbox.environment")
	checkOutput(t, "cbox.environment -> dev", "environment setting not found")

	ctrl.ConfigGet("cbox.default-space")
	checkOutput(t, "cbox.default-space -> default", "default space setting not found")
}

func TestSpaceListingWhenNoSpace(t *testing.T) {
	tty.MockedOutput = ""

	os.RemoveAll("/tmp/.cbox")

	ctrl := controllers.InitController("/tmp")
	ctrl.SpacesDestroy("@default")

	tty.MockedOutput = ""
	ctrl.SpacesList()

	if tty.MockedOutput != "" {
		t.Errorf("there was output while listing empty space list: %s", tty.MockedOutput)
	}
}
