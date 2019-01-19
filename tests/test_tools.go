package tests

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/dplabs/cbox/src/models"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return strings.ToLower(string(b))
}

func findSpaceFile(cboxInstance *models.CBox, space *models.Space) bool {
	spaces := cboxInstance.Spaces

	found := false
	for _, s := range spaces {
		if s.Selector.Namespace == space.Selector.Namespace && s.Label == space.Label {
			found = true
			break
		}
	}
	return found
}

func createSpace(t *testing.T, cboxInstance *models.CBox) *models.Space {
	if cboxInstance == nil {
		t.Fatal("cboxInstance not initialized")
	}

	space := models.Space{
		Label:       randString(8),
		Description: randString(15),
	}
	space.Selector = models.NewSelector(models.TypeUser, "test", space.Label, "")

	err := cboxInstance.SpaceCreate(&space)
	if err != nil {
		t.Fatal(err)
	}

	s, err := cboxInstance.SpaceFind(space.Selector.NamespaceType, space.Selector.Namespace, space.Label)
	if err != nil {
		t.Fatalf("could not find space: %v", err)
	}
	return s
}

func assertSpaceFileExists(t *testing.T, cboxInstance *models.CBox, space *models.Space) {
	found := findSpaceFile(cboxInstance, space)
	if !found {
		t.Fatal("space file could not be found (and should)")
	}
}

func assertSpaceFileNotExists(t *testing.T, cboxInstance *models.CBox, space *models.Space) {
	found := findSpaceFile(cboxInstance, space)
	if found {
		t.Fatal("new space found (and shouldn't)")
	}
}

func assertSliceEqual(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
