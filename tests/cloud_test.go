package tests

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/dpecos/cbox/internal/app/core"
	"github.com/dpecos/cbox/pkg/models"
)

const (
	// jwt for user 'test' in dev
	jwt = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTkxNDgyMjEsImp0aSI6InRlc3QtS0RXU0ciLCJpYXQiOjE1Mzk5NDgyMjEsIm5iZiI6MTUzOTk0ODIyMSwic3ViIjoiLTEiLCJsb2dpbiI6InRlc3QiLCJuYW1lIjoiVGVzdCB1c2VyIn0.w4qpDwWZUjS0NZBmMbYqgg3mE7iucJPpRzAsgSF_936laPBiXe8Lti8r-NvI6jPPQlJCq43JMWg5XersOLRLiJRq4U7HHQdovShcT7U862ZJnWBhJq9famNAJqe7qpuC2BqZWX6bU8QAZhZ_We60_KBsDi7Y2CnK0bWK-MUW8FVgBsGZts-vHxBoon_6W0hFqRL57ncZAS9jua3uGElEW84Ukpgc3ZxFo2oNrrgjFz1WaHxYTMzQx3lOlWFyHEMb6Njslo6nWov-uKcY0eVvOx5mQkLAd33NJk9B0eV8FAXKvn5K2rIIECIfGB6f77teRvQxoN28QNv_OOqKpTAoYA"
)

func TestMain(m *testing.M) {
	if !strings.Contains(core.CloudURL(), "dev") {
		panic("cloud dev environment not set properly")
	}

	os.Exit(m.Run())
}

func TestCloudLogin(t *testing.T) {
	_, login, _, err := core.CloudLogin(jwt)
	if err != nil {
		t.Fatalf("could not login: %v", err)
	}

	if login != "test" {
		t.Errorf("login did not return a test user")
	}
}

func TestSpacePublishingDoesntChangeCreateUpdateDates(t *testing.T) {
	cboxInstance := initializeCBox()

	space := createSpace(t, cboxInstance)
	createCommand(t, space)

	core.Save(cboxInstance)

	_, _, _, err := core.CloudLogin(jwt)
	if err != nil {
		t.Fatalf("could not login: %v", err)
	}

	cloud, err := core.CloudClient()
	if err != nil {
		t.Fatalf("could create cloud client: %v", err)
	}

	err = cloud.SpacePublish(space)
	if err != nil {
		t.Fatalf("could not publish space: %v", err)
	}

	selector, err := models.ParseSelectorForCloudCommand(fmt.Sprintf("@%s:%s", "test", space.Label))
	if err != nil {
		t.Fatalf("could not parse selector for cloud space: %v", err)
	}

	cloudSpace, err := cloud.SpaceRetrieve(selector, nil)
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
