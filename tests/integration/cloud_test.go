package integration_tests

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/dplabs/cbox/src/core"
	"github.com/dplabs/cbox/src/models"
)

const (
	// testUserJWTToken for user 'test' in test env
	testUserJWTToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTkxNDgyMjEsImp0aSI6InRlc3QtS0RXU0ciLCJpYXQiOjE1Mzk5NDgyMjEsIm5iZiI6MTUzOTk0ODIyMSwic3ViIjoiLTEiLCJsb2dpbiI6InRlc3QiLCJuYW1lIjoiVGVzdCB1c2VyIn0.w4qpDwWZUjS0NZBmMbYqgg3mE7iucJPpRzAsgSF_936laPBiXe8Lti8r-NvI6jPPQlJCq43JMWg5XersOLRLiJRq4U7HHQdovShcT7U862ZJnWBhJq9famNAJqe7qpuC2BqZWX6bU8QAZhZ_We60_KBsDi7Y2CnK0bWK-MUW8FVgBsGZts-vHxBoon_6W0hFqRL57ncZAS9jua3uGElEW84Ukpgc3ZxFo2oNrrgjFz1WaHxYTMzQx3lOlWFyHEMb6Njslo6nWov-uKcY0eVvOx5mQkLAd33NJk9B0eV8FAXKvn5K2rIIECIfGB6f77teRvQxoN28QNv_OOqKpTAoYA"
)

var cloud *models.Cloud

func TestMain(m *testing.M) {
	cbox := initializeCBox()
	cloud = cloudConnect(cbox, testUserJWTToken)

	if !strings.Contains(cloud.URL, "test") {
		panic("test setup: cloud test environment not set properly")
	}

	os.Exit(m.Run())
}

func cloudConnect(cbox *models.CBox, jwt string) *models.Cloud {

	cloud := core.CloudClient(cbox)

	_, err := cloud.ServerLogin(jwt)
	if err != nil {
		log.Fatalf("test setup: could not login: %v", err)
	}

	return cloud
}

func TestCloudLogin(t *testing.T) {
	if cloud.Login != "test" {
		t.Errorf("login did not return a test user")
	}
}

func TestSpacePublishingDoesntChangeCreateUpdateDates(t *testing.T) {
	cboxInstance := initializeCBox()

	space := createSpace(t, cboxInstance)
	createCommand(t, space)

	core.Save(cboxInstance)

	err := cloud.SpacePublish(space)
	if err != nil {
		t.Fatalf("could not publish space: %v", err)
	}

	selector, err := models.ParseSelectorForCloud(fmt.Sprintf("@%s:%s", "test", space.Label))
	if err != nil {
		t.Fatalf("could not parse selector for cloud space: %v", err)
	}

	cloudSpace, err := cloud.SpaceFind(selector)
	if err != nil {
		t.Fatalf("could not retrieve space: %v", err)
	}

	if space.CreatedAt != cloudSpace.CreatedAt {
		t.Errorf("creation dates are different between local & cloud spaces: %v", err)
	}

	if space.UpdatedAt != cloudSpace.UpdatedAt {
		t.Errorf("last updated dates are different between local & cloud spaces: %v", err)
	}

	err = cloud.SpaceUnpublish(selector)
	if err != nil {
		t.Fatalf("could not unpublish space: %v", err)
	}

	cloudCommands, err := cloud.CommandList(selector)
	if err != nil {
		t.Fatalf("could not retrieve space commands: %v", err)
	}

	if len(cloudCommands) != 0 {
		t.Errorf("commands left behind in the cloud after deleting their space")
	}
}

func TestPublishingEmptySpace(t *testing.T) {
	cboxInstance := initializeCBox()

	space := createSpace(t, cboxInstance)

	core.Save(cboxInstance)

	err := cloud.SpacePublish(space)
	if err != nil {
		t.Fatalf("could not publish space: %v", err)
	}

	selector, err := models.ParseSelectorForCloud(fmt.Sprintf("@%s:%s", "test", space.Label))
	if err != nil {
		t.Fatalf("could not parse selector for cloud space: %v", err)
	}

	err = cloud.SpaceUnpublish(selector)
	if err != nil {
		t.Errorf("could not unpublish space: %v", err)
	}
}

func TestUnpublishingNonExistingSpace(t *testing.T) {
	initializeCBox()

	selector, err := models.ParseSelectorForCloud(fmt.Sprintf("@%s:%s", "test", "this-space-doesnt-exist"))
	if err != nil {
		t.Fatalf("could not parse selector for cloud space: %v", err)
	}

	err = cloud.SpaceUnpublish(selector)
	if err == nil {
		t.Errorf("did not fail to unpublish a non existing space: %v", err)
	}
}

func TestSpacePublishingDeletesLocallyDeletedCommands(t *testing.T) {
	cboxInstance := initializeCBox()
	space := createSpace(t, cboxInstance)
	command := createCommand(t, space)

	core.Save(cboxInstance)

	err := cloud.SpacePublish(space)
	if err != nil {
		t.Fatalf("could not publish space: %v", err)
	}

	selector, err := models.ParseSelectorForCloud(fmt.Sprintf("@%s:%s", "test", space.Label))
	if err != nil {
		t.Fatalf("could not parse selector for cloud space: %v", err)
	}

	commands, err := cloud.CommandList(selector)
	if err != nil {
		t.Fatalf("could not retrieve commands: %v", err)
	}

	if len(commands) != 1 || commands[0].Label != command.Label {
		t.Errorf("failed to retrieved published commands: '%s'", command.Label)
	}

	space.CommandDelete(command)

	cboxInstance = reloadCBox(cboxInstance)

	err = cloud.SpacePublish(space)
	if err != nil {
		t.Fatalf("could not re-publish space: %v", err)
	}

	commands, err = cloud.CommandList(selector)
	if err != nil {
		t.Fatalf("could not retrieve commands: %v", err)
	}

	err = cloud.SpaceUnpublish(selector)
	if err != nil {
		t.Errorf("could not unpublish space: %v", err)
	}

	if len(commands) != 0 {
		t.Errorf("locally deleted command retrieve after re-publishing: %v", commands)
	}
}
