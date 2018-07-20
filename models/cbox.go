package models

import (
	"fmt"
)

type CBox struct {
	path   string
	Spaces []Space
}

func (cbox *CBox) spaceFindPosition(spaceLabel string) (int, error) {
	for i, space := range cbox.Spaces {
		if space.Label == spaceLabel {
			return i, nil
		}
	}
	return -1, fmt.Errorf("space with label '%s' not found", spaceLabel)
}

func (cbox *CBox) SpaceLabels() []string {
	labels := make(map[string]bool)
	for _, space := range cbox.Spaces {
		if _, ok := labels[space.Label]; !ok {
			labels[space.Label] = true
		}
	}
	result := []string{}
	for k, _ := range labels {
		result = append(result, k)
	}
	return result
}

func (cbox *CBox) SpaceFind(spaceLabel string) *Space {
	pos, err := cbox.spaceFindPosition(spaceLabel)
	if err != nil {
		return nil
	}
	return &cbox.Spaces[pos]
}

func (cbox *CBox) SpaceAdd(space *Space) error {
	if cbox.SpaceFind(space.Label) != nil {
		return fmt.Errorf("space add: space with label '%s' already exists", space.Label)
	}
	cbox.Spaces = append(cbox.Spaces, *space)
	return nil
}
