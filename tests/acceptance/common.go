package acceptance_tests

import (
	"os"
	"strings"
	"testing"

	"github.com/dplabs/cbox/src/controllers"
	"github.com/dplabs/cbox/src/tools/tty"
)

func init() {
	os.Setenv("CBOX_ENV", "test")

	controllers.SkipQuestionsFlag = true
	tty.DisableColors = true
	tty.MockTTY = true
}

func checkOutput(t *testing.T, expected string, msg string) {
	if !strings.Contains(tty.MockedOutput, expected) {
		t.Errorf("%s: %s", msg, tty.MockedOutput)
	}
}
