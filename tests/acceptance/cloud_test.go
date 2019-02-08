package acceptance_tests

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/dplabs/cbox/src/controllers"
	"github.com/dplabs/cbox/src/tools/tty"
)

const (
	// testUserJWTToken for user 'test' in dev
	testUserJWTToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTkxNDgyMjEsImp0aSI6InRlc3QtS0RXU0ciLCJpYXQiOjE1Mzk5NDgyMjEsIm5iZiI6MTUzOTk0ODIyMSwic3ViIjoiLTEiLCJsb2dpbiI6InRlc3QiLCJuYW1lIjoiVGVzdCB1c2VyIn0.w4qpDwWZUjS0NZBmMbYqgg3mE7iucJPpRzAsgSF_936laPBiXe8Lti8r-NvI6jPPQlJCq43JMWg5XersOLRLiJRq4U7HHQdovShcT7U862ZJnWBhJq9famNAJqe7qpuC2BqZWX6bU8QAZhZ_We60_KBsDi7Y2CnK0bWK-MUW8FVgBsGZts-vHxBoon_6W0hFqRL57ncZAS9jua3uGElEW84Ukpgc3ZxFo2oNrrgjFz1WaHxYTMzQx3lOlWFyHEMb6Njslo6nWov-uKcY0eVvOx5mQkLAd33NJk9B0eV8FAXKvn5K2rIIECIfGB6f77teRvQxoN28QNv_OOqKpTAoYA"
)

func TestLogInAndOutToCloud(t *testing.T) {
	dir, _ := ioutil.TempDir("", "cbox")
	defer os.RemoveAll(dir)
	ctrl := controllers.InitController(dir)

	ctrl.ConfigSet("cbox.environment", "dev")

	tty.MockedOutput = ""
	tty.MockedInput = []string{testUserJWTToken}
	ctrl.CloudLogin()
	checkOutput(t, "Hi Test user!", "failed to login")

	tty.MockedOutput = ""
	ctrl.CloudLogout()
	checkOutput(t, "Successfully logged out from cbox cloud. See you back soon!", "failed to logout")
}

func TestLogPublishingAndUnpublishingToCloud(t *testing.T) {
	dir, _ := ioutil.TempDir("", "cbox")
	defer os.RemoveAll(dir)
	ctrl := controllers.InitController(dir)

	ctrl.ConfigSet("cbox.environment", "dev")

	tty.MockedInput = []string{testUserJWTToken}
	ctrl.CloudLogin()

	tty.MockedInput = []string{"test-command", "This is a test command", "URL", "CODE", "test-tag"}
	ctrl.CommandAdd(nil)

	tty.MockedOutput = ""
	ctrl.CloudSpacePublish("@default")
	checkOutput(t, "Space published successfully!", "failed to publish space")

	tty.MockedOutput = ""
	ctrl.CloudSpaceInfo("@test:default")
	checkOutput(t, "@test:default - Default space to store commands", "failed to retrieve published space")

	tty.MockedOutput = ""
	ctrl.CloudSpaceUnpublish("@test:default")
	checkOutput(t, "Space unpublished successfully!", "failed to unpublish space")

	// TODO: figure out how to catch a log.Fatalf to check that it was actually invoked
	// tty.MockedOutput = ""
	// ctrl.CloudSpaceInfo("@test:default")
	// checkOutput(t, "Space '@test:default' not found: rest: request failed with '404 Not Found'", "did retrieve space info after deleting it")
}
