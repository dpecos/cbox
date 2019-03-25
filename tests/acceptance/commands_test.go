package acceptance_tests

import (
	"os"
	"strings"
	"testing"

	"github.com/dplabs/cbox/src/controllers"
	"github.com/dplabs/cbox/src/tools/tty"
	"github.com/dplabs/cbox/tests"
)

func TestAddDeleteCommand(t *testing.T) {
	ctrl, dir := tests.InitController()
	defer os.RemoveAll(dir)

	tty.MockedOutput = ""
	tty.MockedInput = []string{"test-command", "This is a test command", "url", "CODE", "test-tag"}
	ctrl.CommandAdd(nil)
	tests.AssertOutputContains(t, "Selector: test-command@default", "failed to add a command")
	tests.AssertOutputContains(t, "url", "could not display the command (url) in the default space")
	tests.AssertOutputContains(t, "CODE", "could not display the command (code) in the default space")

	tty.MockedOutput = ""
	ctrl.CommandList(nil)
	tests.AssertOutputContains(t, "test-command@default - This is a test command (test-tag)", "could not find the command in the default space")

	tty.MockedOutput = ""
	ctrl.CommandDelete("test-command@default")

	tty.MockedOutput = ""
	ctrl.CommandList(nil)
	if strings.Contains(tty.MockedOutput, "test-command@default") {
		t.Errorf("failed to delete the command from space: %s", tty.MockedOutput)
	}
}

func TestEditViewCommand(t *testing.T) {
	ctrl, dir := tests.InitController()
	defer os.RemoveAll(dir)

	tty.MockedInput = []string{"test-command", "This is a test command", "url", "CODE", "test-tag"}
	ctrl.CommandAdd(nil)

	tty.MockedOutput = ""
	ctrl.CommandView("test-command@default")
	tests.AssertOutputContains(t, "Selector: test-command@default", "could not display the command in the default space")
	tests.AssertOutputContains(t, "url", "could not display the command (url) in the default space")
	tests.AssertOutputContains(t, "CODE", "could not display the command (code) in the default space")

	tty.MockedOutput = ""
	tty.MockedInput = []string{"test-command-edited", "This is a test command - edited", "url-edited", "CODE-edited"}
	ctrl.CommandEdit("test-command@default")
	tests.AssertOutputContains(t, "Selector: test-command-edited@default", "failed to edit a command")
	tests.AssertOutputContains(t, "url-edited", "failed to edit the command (url) in the default space")
	tests.AssertOutputContains(t, "CODE-edited", "failed to edit the command (code) in the default space")

	tty.MockedOutput = ""
	ctrl.CommandView("test-command-edited@default")
	tests.AssertOutputContains(t, "Selector: test-command-edited@default", "could not display the command in the default space")
	tests.AssertOutputContains(t, "url-edited", "could not display the command (url) in the default space")
	tests.AssertOutputContains(t, "CODE-edited", "could not display the command (code) in the default space")
}

func TestTagCommand(t *testing.T) {
	ctrl, dir := tests.InitController()
	defer os.RemoveAll(dir)

	tty.MockedInput = []string{"test-command", "This is a test command", "url", "CODE", "test-tag"}
	ctrl.CommandAdd(nil)

	tty.MockedOutput = ""
	ctrl.TagsAdd("test-command@default", "t1", "t2", "t3")
	tests.AssertOutputContains(t, "Tags: test-tag, t1, t2, t3", "failed to add tags to command")

	tty.MockedOutput = ""
	ctrl.TagsRemove("test-command@default", "t1", "t2")
	tests.AssertOutputContains(t, "Tags: test-tag, t3", "failed to remove tags from a command")

	tty.MockedOutput = ""
	ctrl.TagsList(nil)
	tests.AssertOutputContains(t, "* t3", "could not list tags (t3) in the default space")
	tests.AssertOutputContains(t, "* test-tag", "could not list tags (test-tag) in the default space")
	if strings.Contains(tty.MockedOutput, "t1") {
		t.Errorf("failed to delete tag t1 from command: %s", tty.MockedOutput)
	}
	if strings.Contains(tty.MockedOutput, "t2") {
		t.Errorf("failed to delete tag t2 from command: %s", tty.MockedOutput)
	}

	tty.MockedOutput = ""
	ctrl.TagsDelete("t3@default")
	tests.AssertOutputContains(t, "Tag 't3' successfully deleted from space 'default'", "could not delete tag t3 in the default space")

	tty.MockedOutput = ""
	ctrl.CommandView("test-command@default")
	tests.AssertOutputContains(t, "Selector: test-command@default", "could not display the command in the default space")
}

func TestCopyCommand(t *testing.T) {
	ctrl, dir := tests.InitController()
	defer os.RemoveAll(dir)

	tty.MockedInput = []string{"test-command", "This is a test command", "url", "CODE", "test-tag"}
	ctrl.CommandAdd(nil)

	tty.MockedOutput = ""
	tty.MockedInput = []string{"test-space", "This is a test space"}
	ctrl.SpacesCreate()

	// specifying both origin and target
	targetSpace := "@test-space"
	tty.MockedOutput = ""
	ctrl.CommandCopy("test-command@default", &targetSpace)
	tests.AssertOutputContains(t, "Command copied successfully!", "could not copy command to @test-space")

	tty.MockedOutput = ""
	ctrl.CommandView("test-command@test-space")
	tests.AssertOutputContains(t, "Selector: test-command@test-space", "could not display the command in @test-space")

	tty.MockedOutput = ""
	ctrl.CommandView("test-command@default")
	tests.AssertOutputContains(t, "Selector: test-command@default", "could not display the command in @default")

	tty.MockedOutput = ""
	ctrl.CommandDelete("test-command@test-space")

	// specifying both origin (without space) and target
	tty.MockedOutput = ""
	ctrl.CommandCopy("test-command", &targetSpace)
	tests.AssertOutputContains(t, "Command copied successfully!", "could not copy command to @test-space (origin space not specified)")

	tty.MockedOutput = ""
	ctrl.CommandView("test-command@test-space")
	tests.AssertOutputContains(t, "Selector: test-command@test-space", "could not display the command in @test-space (origin space not specified)")

	tty.MockedOutput = ""
	ctrl.CommandView("test-command@default")
	tests.AssertOutputContains(t, "Selector: test-command@default", "could not display the command in @default (origin space not specified)")

	tty.MockedOutput = ""
	ctrl.CommandDelete("test-command@default")

	// specifying only origin (without space)
	tty.MockedOutput = ""
	ctrl.CommandCopy("test-command@test-space", nil)
	tests.AssertOutputContains(t, "Command copied successfully!", "could not copy command to @default (default as target space)")

	tty.MockedOutput = ""
	ctrl.CommandView("test-command@default")
	tests.AssertOutputContains(t, "Selector: test-command@default", "could not display the command in @default (default as target space)")

	tty.MockedOutput = ""
	ctrl.CommandView("test-command@test-space")
	tests.AssertOutputContains(t, "Selector: test-command@test-space", "could not display the command in @test-space (default as target space)")
}

func TestCopyCommandClashing(t *testing.T) {
	ctrl, dir := tests.InitController()
	defer os.RemoveAll(dir)

	tty.MockedInput = []string{"test-command", "This is a test command", "url", "CODE", "test-tag"}
	ctrl.CommandAdd(nil)

	tty.MockedOutput = ""
	tty.MockedInput = []string{"test-space", "This is a test space"}
	ctrl.SpacesCreate()

	targetSpace := "@test-space"
	tty.MockedOutput = ""
	ctrl.CommandCopy("test-command@default", &targetSpace)

	controllers.ForceFlag = true
	tty.MockedOutput = ""
	ctrl.CommandCopy("test-command@default", &targetSpace)
	tests.AssertOutputContains(t, "Command copied successfully!", "could not copy command to @test-space")

	// TODO: check for error thrown
	// controllers.ForceFlag = false
	// tty.MockedOutput = ""
	// ctrl.CommandCopy("test-command@default", &targetSpace)
	// tests.AssertOutputContains(t, "Command copied successfully!", "could not copy command to @test-space")
}
