package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/dpecos/cbox/models"
)

const SPACES_PATH = "spaces"
const DEFAULT_SPACE_NAME = "default"
const DEFAULT_SPACE_DESCRIPTION = "Default space to store commands"

func resolveSpaceFile(spaceName string) string {
	spacePath := path.Join(SPACES_PATH, spaceName+".json")
	return resolveInCboxDir(spacePath)
}

func SpaceList() []*models.Space {

	spaces := []*models.Space{}
	files, err := ioutil.ReadDir(resolveInCboxDir(SPACES_PATH))
	if err != nil {
		log.Fatalf("repository: could not read spaces: %v", err)
	}
	for _, f := range files {
		filename := f.Name()
		extension := filepath.Ext(filename)
		if extension == ".json" {
			name := filename[0 : len(filename)-len(extension)]
			spaces = append(spaces, SpaceLoad(name))
		}
	}
	return spaces
}

func SpaceStore(space *models.Space) {
	raw, err := json.MarshalIndent(space, "", "  ")
	if err != nil {
		log.Fatalf("repository: store space %s: could not generate JSON: %v", space.Name, err)
	}

	file := resolveSpaceFile(space.Name)
	err = ioutil.WriteFile(file, raw, 0644)
	if err != nil {
		log.Fatalf("repository: store space %s: could not write JSON file (%s): %v", space.Name, file, err)
	}
}

func SpaceDelete(space *models.Space) {
	err := os.Remove(resolveSpaceFile(space.Name))
	if err != nil {
		log.Fatalf("repository: delete space %s: %v", space.Name, err)
	}
}

func SpaceLoad(name string) *models.Space {
	raw, err := ioutil.ReadFile(resolveSpaceFile(name))
	if err != nil {
		log.Fatalf("repository: load space %s: could not read file: %v", name, err)
	}

	var space models.Space
	err = json.Unmarshal(raw, &space)
	if err != nil {
		log.Fatalf("repository: load space %s: could not parse JSON file: %v", name, err)
	}

	return &space
}
