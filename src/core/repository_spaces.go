package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/dplabs/cbox/src/models"
)

const (
	filenameSeparatorUser         = ":"
	filenameSeparatorOrganization = "="
)

func resolveSpaceFile(namespace string, label string) string {
	filename := label
	if namespace != "" {
		namespace = strings.Replace(namespace, "/", filenameSeparatorOrganization, -1)
		filename = fmt.Sprintf("%s%s", namespace, label)
	}
	filename = filename + ".json"
	return resolveInCboxDir(path.Join(pathSpaces, filename))
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
			namespace := ""
			label := filename[0 : len(filename)-len(extension)]
			if strings.Contains(label, filenameSeparatorUser) {
				parts := strings.Split(label, filenameSeparatorUser)
				namespace = models.SUser(parts[0])
				label = parts[1]
			} else if strings.Contains(label, filenameSeparatorOrganization) {
				parts := strings.Split(label, filenameSeparatorOrganization)
				namespace = models.SOrganization(parts[0])
				label = parts[1]
			}
			spaces = append(spaces, spaceLoadFile(namespace, label))
		}
	}
	return spaces
}

func spaceLoadFile(namespace string, label string) *models.Space {
	spacePath := resolveSpaceFile(namespace, label)

	raw, err := ioutil.ReadFile(spacePath)
	if err != nil {
		log.Fatalf("repository: load space '%s%s': could not read file '%s': %v", namespace, label, spacePath, err)
	}

	var space models.Space
	err = json.Unmarshal(raw, &space)

	space.Namespace = namespace

	if err != nil {
		log.Fatalf("repository: load space '%s%s': could not parse JSON file: %v", namespace, label, err)
	}

	if space.Entries == nil {
		space.Entries = []*models.Command{}
	}

	for _, command := range space.Entries {
		command.Space = &space
	}

	return &space
}

func spaceStoreFile(space *models.Space) {

	namespace := space.Namespace
	space.Namespace = models.SNamespace(namespace)

	raw, err := json.MarshalIndent(space, "", "  ")
	if err != nil {
		log.Fatalf("repository: store space '%s': could not generate JSON: %v", space.String(), err)
	}

	space.Namespace = namespace

	file := resolveSpaceFile(space.Namespace, space.Label)
	err = ioutil.WriteFile(file, raw, 0644)
	if err != nil {
		log.Fatalf("repository: store space '%s': could not write JSON file (%s): %v", space.String(), file, err)
	}
}

func spaceDeleteFile(space *models.Space) {
	file := resolveSpaceFile(space.Namespace, space.Label)
	err := os.Remove(file)
	if err != nil {
		log.Fatalf("repository: delete space '%s': %v", space.String(), err)
	}
}
