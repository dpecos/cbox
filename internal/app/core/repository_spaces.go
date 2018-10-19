package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/dpecos/cbox/pkg/models"
)

func resolveSpaceFile(spaceName string) string {
	spacePath := path.Join(pathSpaces, spaceName+".json")
	return resolveInCboxDir(spacePath)
}

func spacesLoad() []*models.Space {
	spaces := []*models.Space{}
	files, err := ioutil.ReadDir(resolveInCboxDir(pathSpaces))
	if err != nil {
		log.Fatalf("repository: could not read spaces: %v", err)
	}
	for _, f := range files {
		filename := f.Name()
		extension := filepath.Ext(filename)
		if extension == ".json" {
			name := filename[0 : len(filename)-len(extension)]
			spaces = append(spaces, spaceLoadFile(name))
		}
	}
	return spaces
}

func spaceLoadFile(label string) *models.Space {
	raw, err := ioutil.ReadFile(resolveSpaceFile(label))
	if err != nil {
		log.Fatalf("repository: load space '%s': could not read file: %v", label, err)
	}

	var space models.Space
	err = json.Unmarshal(raw, &space)

	if err != nil {
		log.Fatalf("repository: load space '%s': could not parse JSON file: %v", label, err)
	}

	if space.Entries == nil {
		space.Entries = []models.Command{}
	}

	return &space
}

func spaceStoreFile(space *models.Space) {
	raw, err := json.MarshalIndent(space, "", "  ")
	if err != nil {
		log.Fatalf("repository: store space '%s': could not generate JSON: %v", space.Label, err)
	}

	file := resolveSpaceFile(space.Label)
	err = ioutil.WriteFile(file, raw, 0644)
	if err != nil {
		log.Fatalf("repository: store space '%s': could not write JSON file (%s): %v", space.Label, file, err)
	}
}

func spaceDeleteFile(space *models.Space) {
	file := resolveSpaceFile(space.Label)
	err := os.Remove(file)
	if err != nil {
		log.Fatalf("repository: delete space '%s': %v", space.Label, err)
	}
}
