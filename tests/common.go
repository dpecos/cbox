package tests

import (
	"math/rand"
	"testing"
	"time"

	"github.com/dpecos/cbox/core"
	"github.com/dpecos/cbox/models"
)

var (
	cbox *models.CBox
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func findSpaceFile(space *models.Space) bool {
	spaces := core.SpaceList()

	found := false
	for _, s := range spaces {
		if s.ID == space.ID {
			found = true
			break
		}
	}
	return found
}

func assertSpaceFileExists(t *testing.T, space *models.Space) {
	found := findSpaceFile(space)
	if !found {
		t.Fatal("space file could not be found (and should)")
	}
}
func assertSpaceFileNotExists(t *testing.T, space *models.Space) {
	found := findSpaceFile(space)
	if found {
		t.Fatal("new space found (and shouldn't)")
	}
}
