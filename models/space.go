package models

import "github.com/dpecos/cbox/tools/console"

type Space struct {
	Name    string    `json:"name"`
	Title   string    `json:"title"`
	Entries []Command `json:"entries"`
}

func commandPresentInSapce(space *Space, command *Command) bool {
	for _, cmd := range space.Entries {
		if command.ID == cmd.ID {
			return true
		}
	}
	return false
}

func (space *Space) CommandAdd(command *Command) {
	for commandPresentInSapce(space, command) {
		console.PrintError("ID already found in space. Try a different one")
		command.ID = console.ReadString("ID")
	}
	space.Entries = append(space.Entries, *command)
}

func (space *Space) CommandList(tag string) []Command {
	if tag == "" {
		return space.Entries
	}

	var result []Command
	for _, command := range space.Entries {
		for _, t := range command.Tags {
			if tag == t {
				result = append(result, command)
			}
		}
	}
	return result
}
