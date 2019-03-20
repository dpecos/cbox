package acceptance_tests

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/dplabs/cbox/src/controllers"
	"github.com/dplabs/cbox/src/tools/tty"
)

func TestAddDeleteCommand(t *testing.T) {
	dir, _ := ioutil.TempDir("", "cbox")
	defer os.RemoveAll(dir)
	ctrl := controllers.InitController(dir)

	tty.MockedOutput = ""
	tty.MockedInput = []string{"test-command", "This is a test command", "url", "CODE", "test-tag"}
	ctrl.CommandAdd(nil)
	checkOutput(t, "Selector: test-command@default", "failed to add a command")
	checkOutput(t, "url", "could not display the command (url) in the default space")
	checkOutput(t, "CODE", "could not display the command (code) in the default space")

	tty.MockedOutput = ""
	ctrl.CommandList(nil)
	checkOutput(t, "test-command@default - This is a test command (test-tag)", "could not find the command in the default space")

	tty.MockedOutput = ""
	ctrl.CommandDelete("test-command@default")

	tty.MockedOutput = ""
	ctrl.CommandList(nil)
	if strings.Contains(tty.MockedOutput, "test-command@default") {
		t.Errorf("failed to delete the command from space: %s", tty.MockedOutput)
	}
}

func TestEditViewCommand(t *testing.T) {
	dir, _ := ioutil.TempDir("", "cbox")
	defer os.RemoveAll(dir)
	ctrl := controllers.InitController(dir)

	tty.MockedInput = []string{"test-command", "This is a test command", "url", "CODE", "test-tag"}
	ctrl.CommandAdd(nil)

	tty.MockedOutput = ""
	ctrl.CommandView("test-command@default")
	checkOutput(t, "Selector: test-command@default", "could not display the command in the default space")
	checkOutput(t, "url", "could not display the command (url) in the default space")
	checkOutput(t, "CODE", "could not display the command (code) in the default space")

	tty.MockedOutput = ""
	tty.MockedInput = []string{"test-command-edited", "This is a test command - edited", "url-edited", "CODE-edited"}
	ctrl.CommandEdit("test-command@default")
	checkOutput(t, "Selector: test-command-edited@default", "failed to edit a command")
	checkOutput(t, "url-edited", "failed to edit the command (url) in the default space")
	checkOutput(t, "CODE-edited", "failed to edit the command (code) in the default space")

	tty.MockedOutput = ""
	ctrl.CommandView("test-command-edited@default")
	checkOutput(t, "Selector: test-command-edited@default", "could not display the command in the default space")
	checkOutput(t, "url-edited", "could not display the command (url) in the default space")
	checkOutput(t, "CODE-edited", "could not display the command (code) in the default space")
}

func TestTagCommand(t *testing.T) {
	dir, _ := ioutil.TempDir("", "cbox")
	defer os.RemoveAll(dir)
	ctrl := controllers.InitController(dir)

	tty.MockedInput = []string{"test-command", "This is a test command", "url", "CODE", "test-tag"}
	ctrl.CommandAdd(nil)

	tty.MockedOutput = ""
	ctrl.TagsAdd("test-command@default", "t1", "t2", "t3")
	checkOutput(t, "Tags: test-tag, t1, t2, t3", "failed to add tags to command")

	tty.MockedOutput = ""
	ctrl.TagsRemove("test-command@default", "t1", "t2")
	checkOutput(t, "Tags: test-tag, t3", "failed to remove tags from a command")

	tty.MockedOutput = ""
	ctrl.TagsList(nil)
	checkOutput(t, "* t3", "could not list tags (t3) in the default space")
	checkOutput(t, "* test-tag", "could not list tags (test-tag) in the default space")
	if strings.Contains(tty.MockedOutput, "t1") {
		t.Errorf("failed to delete tag t1 from command: %s", tty.MockedOutput)
	}
	if strings.Contains(tty.MockedOutput, "t2") {
		t.Errorf("failed to delete tag t2 from command: %s", tty.MockedOutput)
	}

	tty.MockedOutput = ""
	ctrl.TagsDelete("t3@default")
	checkOutput(t, "Tag 't3' successfully deleted from space 'default'", "could not delete tag t3 in the default space")

	tty.MockedOutput = ""
	ctrl.CommandView("test-command@default")
	checkOutput(t, "Selector: test-command@default", "could not display the command in the default space")
}
