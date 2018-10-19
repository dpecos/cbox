package tests

import (
	"log"
	"math/rand"
	"strings"
	"testing"

	"github.com/dpecos/cbox/pkg/models"
	"github.com/gofrs/uuid"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
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
		if s.ID == space.ID {
			found = true
			break
		}
	}
	return found
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

func createSpace(t *testing.T, cboxInstance *models.CBox) *models.Space {
	if cboxInstance == nil {
		t.Fatal("cboxInstance not initialized")
	}

	id, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("could not generate a new ID: %v", err)
	}
	space := models.Space{
		Label:       randString(8),
		Description: randString(15),
	}
	space.ID = id

	err = cboxInstance.SpaceCreate(&space)

	if err != nil {
		t.Error(err)
	}

	s, err := cboxInstance.SpaceFind(space.ID.String())
	if err != nil {
		t.Fatalf("could not find space: %v", err)
	}
	return s
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
