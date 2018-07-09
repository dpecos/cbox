package models

import (
	"fmt"
	"log"
	"strings"

	"github.com/dpecos/cbox/tools/console"
)

type CBox struct {
	path   string
	Spaces []Space
}

func (cbox *CBox) spaceFindPosition(spaceName string) (int, error) {
	for i, space := range cbox.Spaces {
		if space.Name == spaceName {
			return i, nil
		}
	}
	return -1, fmt.Errorf("space with name '%s' not found", spaceName)
}

func (cbox *CBox) spaceInCbox(spaceName string) bool {
	pos, err := cbox.spaceFindPosition(spaceName)
	return err == nil && pos != -1
}

func (cbox *CBox) findUniqueSpaceNames() []string {
	names := make(map[string]bool)
	for _, space := range cbox.Spaces {
		if _, ok := names[space.Name]; !ok {
			names[space.Name] = true
		}
	}
	result := []string{}
	for k, _ := range names {
		result = append(result, k)
	}
	return result
}

func (cbox *CBox) SpaceFind(spaceName string) *Space {
	pos, err := cbox.spaceFindPosition(spaceName)
	if err != nil {
		log.Fatalf("space find: %v", err)
	}
	return &cbox.Spaces[pos]
}

func (cbox *CBox) SpaceAdd(space *Space) {
	cbox.Spaces = append(cbox.Spaces, *space)
}

func (cbox *CBox) SpaceCreate(space *Space) {
	for cbox.spaceInCbox(space.Name) {
		console.PrintError("Space already found in your cbox. Try a different one")
		space.Name = strings.ToLower(console.ReadString("Name"))
	}

	cbox.SpaceAdd(space)
}

func (cbox *CBox) SpaceEdit(space *Space, previousName string) {
	if space.Name != previousName {

		for len(cbox.findUniqueSpaceNames()) != len(cbox.Spaces) {
			console.PrintError("Name already found in your cbox. Try a different one")
			space.Name = strings.ToLower(console.ReadString("Name"))
		}

	}
}
