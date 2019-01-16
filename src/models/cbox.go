package models

import (
	"fmt"
)

type CBox struct {
	path   string
	Spaces []*Space
}

func (cbox *CBox) spaceFindPositionByLabel(namespace string, label string) (int, error) {
	if label == "" {
		return -1, fmt.Errorf("could not search by empty label")
	}
	for i, space := range cbox.Spaces {
		if space.Namespace == namespace && space.Label == label {
			return i, nil
		}
	}
	return -1, fmt.Errorf("space '%s%s' not found", namespace, label)
}

func (cbox *CBox) SpaceLabels() []string {
	labels := make(map[string]bool)
	for _, space := range cbox.Spaces {
		if _, ok := labels[space.String()]; !ok {
			labels[space.String()] = true
		}
	}
	result := []string{}
	for k, _ := range labels {
		result = append(result, k)
	}
	return result
}

func (cbox *CBox) SpaceFind(namespace string, label string) (*Space, error) {
	pos, err := cbox.spaceFindPositionByLabel(namespace, label)
	if err != nil {
		return nil, fmt.Errorf("find space: %v", err)
	}
	return cbox.Spaces[pos], nil
}

func (cbox *CBox) SpaceCreate(space *Space) error {
	s, err := cbox.SpaceFind(space.Namespace, space.Label)
	if err == nil && s != nil {
		return fmt.Errorf("space create: space '%s' already exists", space.String())
	}
	if space.Entries == nil {
		space.Entries = []*Command{}
	}
	if space.CreatedAt == NilUnixTime {
		now := UnixTimeNow()
		space.CreatedAt = now
		space.UpdatedAt = now
	}
	cbox.Spaces = append(cbox.Spaces, space)
	return nil
}

func (cbox *CBox) SpaceDestroy(space *Space) error {
	pos, err := cbox.spaceFindPositionByLabel(space.Namespace, space.Label)
	if err != nil {
		return fmt.Errorf("space destroy: could not found space '%s'", space.String())
	}

	cbox.Spaces = append(cbox.Spaces[:pos], cbox.Spaces[pos+1:]...)

	return nil
}

func (cbox *CBox) SpaceEdit(space *Space, previousNamespace string, previousLabel string) error {

	if (space.Namespace != previousNamespace || space.Label != previousLabel) && len(cbox.SpaceLabels()) != len(cbox.Spaces) {
		return fmt.Errorf("space edit: duplicate namespace/label '%s", space.String())
	}

	space.UpdatedAt = UnixTimeNow()

	return nil
}
