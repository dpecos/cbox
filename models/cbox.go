package models

import (
	"log"

	"github.com/dpecos/cbox/tools/console"
)

type CBox struct {
	path   string
	Spaces []Space
}

func (cbox *CBox) SpaceFind(spaceName string) *Space {
	for i, space := range cbox.Spaces {
		if space.Name == spaceName {
			return &cbox.Spaces[i]
		}
	}
	log.Fatalf("Could not find space with name %s", spaceName)
	return nil
}

func (cbox *CBox) SpaceAdd(space *Space) {
	cbox.Spaces = append(cbox.Spaces, *space)
}

func (cbox *CBox) spaceInCbox(space *Space) bool {
	for _, s := range cbox.Spaces {
		if s.Name == space.Name {
			return true
		}
	}
	return false
}

func (cbox *CBox) SpaceCreate(space *Space) {
	for cbox.spaceInCbox(space) {
		console.PrintError("Space already found in your cbox. Try a different one")
		space.Name = console.ReadString("Name")
	}

	cbox.SpaceAdd(space)
}

func (cbox *CBox) SpaceDelete(space Space) {
}
