package acceptance_tests

import (
	"os"
	"strings"
	"testing"

	"github.com/dplabs/cbox/src/tools/tty"
)

func init() {
	os.Setenv("CBOX_ENV", "test")

	tty.DisableColors = true
	tty.MockTTY = true
	tty.SkipQuestions = true
}

func checkOutput(t *testing.T, expected string, msg string) {
	if !strings.Contains(tty.MockedOutput, expected) {
		t.Errorf("%s: %s", msg, tty.MockedOutput)
	}
}
