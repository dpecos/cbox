package acceptance_tests

import (
	"os"
	"testing"

	"github.com/dplabs/cbox/src/controllers"
	"github.com/dplabs/cbox/src/tools/tty"
	"github.com/dplabs/cbox/tests"
)

const (
	// testUserJWTToken for user 'test' in test env
	testUserJWTToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTkxNDgyMjEsImp0aSI6InRlc3QtS0RXU0ciLCJpYXQiOjE1Mzk5NDgyMjEsIm5iZiI6MTUzOTk0ODIyMSwic3ViIjoiLTEiLCJsb2dpbiI6InRlc3QiLCJuYW1lIjoiVGVzdCB1c2VyIn0.w4qpDwWZUjS0NZBmMbYqgg3mE7iucJPpRzAsgSF_936laPBiXe8Lti8r-NvI6jPPQlJCq43JMWg5XersOLRLiJRq4U7HHQdovShcT7U862ZJnWBhJq9famNAJqe7qpuC2BqZWX6bU8QAZhZ_We60_KBsDi7Y2CnK0bWK-MUW8FVgBsGZts-vHxBoon_6W0hFqRL57ncZAS9jua3uGElEW84Ukpgc3ZxFo2oNrrgjFz1WaHxYTMzQx3lOlWFyHEMb6Njslo6nWov-uKcY0eVvOx5mQkLAd33NJk9B0eV8FAXKvn5K2rIIECIfGB6f77teRvQxoN28QNv_OOqKpTAoYA"
)

func TestLogInAndLogOutToCloud(t *testing.T) {
	ctrl, dir := tests.InitController()
	defer os.RemoveAll(dir)

	tty.MockedOutput = ""
	tty.MockedInput = []string{testUserJWTToken}
	ctrl.CloudLogin()
	tests.AssertOutputContains(t, "Hi Test user!", "failed to login")

	tty.MockedOutput = ""
	ctrl.CloudLogout()
	tests.AssertOutputContains(t, "Successfully logged out from cbox cloud. See you back soon!", "failed to logout")
}

func TestPublishingAndUnpublishingToCloud(t *testing.T) {
	ctrl, dir := tests.InitController()
	defer os.RemoveAll(dir)

	tty.MockedInput = []string{testUserJWTToken}
	ctrl.CloudLogin()

	tty.MockedInput = []string{"test-command", "This is a test command", "URL", "CODE", "test-tag"}
	ctrl.CommandAdd(nil)

	tty.MockedOutput = ""
	ctrl.CloudSpacePublish("@default")
	tests.AssertOutputContains(t, "Space published successfully!", "failed to publish space")

	tty.MockedOutput = ""
	ctrl.CloudSpaceInfo("@test:default")
	tests.AssertOutputContains(t, "@test:default - Default space to store commands", "failed to retrieve published space")

	tty.MockedOutput = ""
	ctrl.CloudSpaceUnpublish("@test:default")
	tests.AssertOutputContains(t, "Space unpublished successfully!", "failed to unpublish space")

	// TODO: figure out how to catch a log.Fatalf to check that it was actually invoked
	// tty.MockedOutput = ""
	// ctrl.CloudSpaceInfo("@test:default")
	// tests.AssertOutputContains(t, "Space '@test:default' not found: rest: request failed with '404 Not Found'", "did retrieve space info after deleting it")
}

func TestViewingCloudCommands(t *testing.T) {
	ctrl, dir := tests.InitController()
	defer os.RemoveAll(dir)

	tty.MockedInput = []string{testUserJWTToken}
	ctrl.CloudLogin()

	tty.MockedInput = []string{"test-command", "This is a test command", "URL", "CODE", "test-tag"}
	ctrl.CommandAdd(nil)

	tty.MockedOutput = ""
	ctrl.CloudSpacePublish("@default")
	tests.AssertOutputContains(t, "Space published successfully!", "failed to publish space")

	tty.MockedOutput = ""
	ctrl.CloudCommandView("test-command@test:default")
	tests.AssertOutputContains(t, "Selector: test-command@test:default", "failed to retrieve published command")

	tty.MockedOutput = ""
	ctrl.CloudSpaceUnpublish("@test:default")
	tests.AssertOutputContains(t, "Space unpublished successfully!", "failed to unpublish space")
}

func TestCopyingCloudCommands(t *testing.T) {
	ctrl, dir := tests.InitController()
	defer os.RemoveAll(dir)

	tty.MockedInput = []string{testUserJWTToken}
	ctrl.CloudLogin()

	tty.MockedInput = []string{"test-command", "This is a test command", "URL", "CODE", "test-tag"}
	ctrl.CommandAdd(nil)

	tty.MockedOutput = ""
	ctrl.CloudSpacePublish("@default")

	// clone a remote space
	ctrl.SpacesDestroy("@default")

	tty.MockedOutput = ""
	ctrl.CloudCopy("@test:default", nil)
	tests.AssertOutputContains(t, "Space cloned successfully into '@default'!", "failed to clone space")

	// clone with clashing space
	tty.MockedInput = []string{"clashed"}
	tty.MockedOutput = ""
	ctrl.CloudCopy("@test:default", nil)
	tests.AssertOutputContains(t, "Space cloned successfully into '@clashed'!", "failed to clone space")

	// copy remote commands overwriting local ones
	controllers.ForceFlag = true

	tty.MockedOutput = ""
	targetSpace := "@default"
	ctrl.CloudCopy("@test:default", &targetSpace)
	tests.AssertOutputContains(t, "Commands copied successfully into '@default'!", "failed to clone space (clashing)")

	// copy single remote command overwriting local ones not specifying target space == @default
	tty.MockedOutput = ""
	ctrl.CloudCopy("test-command@test:default", nil)
	tests.AssertOutputContains(t, "Commands copied successfully into '@default'!", "failed to clone space (clashing)")

	// cleanup
	tty.MockedOutput = ""
	ctrl.CloudSpaceUnpublish("@test:default")
	tests.AssertOutputContains(t, "Space unpublished successfully!", "failed to unpublish space")
}
