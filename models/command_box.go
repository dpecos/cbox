package models

type CommandBox struct {
	path   string
	Spaces []Space
}

func (cbox *CommandBox) SpaceFind(spaceName string) *Space {
	for _, space := range cbox.Spaces {
		if space.Name == spaceName {
			return &space
		}
	}
	return nil
}

func (cbox *CommandBox) SpaceAdd(space Space) Space {
	cbox.Spaces = append(cbox.Spaces, space)
	return space
}

func (cbox *CommandBox) SpaceCreate(spaceName string) Space {
	var space Space
	return space
}

func (cbox *CommandBox) SpaceDelete(space Space) {
}
