package models

import (
	"fmt"
	"log"
)

type Space struct {
	Meta
	Label       string     `json:"label"`
	Description string     `json:"description"`
	Entries     []*Command `json:"entries" dynamodbav:"-"`
}

func commandPresentInSpace(space *Space, commandLabel string) bool {
	_, err := space.commandFindPositionByLabel(commandLabel)
	return err == nil
}

func (space *Space) CommandAdd(command *Command) error {
	if commandPresentInSpace(space, command.Label) {
		return fmt.Errorf("add command: label '%s' already in use", command.Label)
	}
	now := UnixTimeNow()
	if command.CreatedAt == NilUnixTime {
		command.CreatedAt = now
		command.UpdatedAt = now
	}
	space.UpdatedAt = now
	space.Entries = append(space.Entries, command)

	return nil
}

func (space *Space) CommandEdit(command *Command, previousLabel string) error {
	if command.Label != previousLabel {
		newLabel := command.Label
		command.Label = previousLabel
		if commandPresentInSpace(space, newLabel) {
			command.Label = newLabel
			return fmt.Errorf("edit command: label '%s' already in use", newLabel)
		}
		command.Label = newLabel
	}
	now := UnixTimeNow()
	command.UpdatedAt = now
	space.UpdatedAt = now

	return nil
}

func (space *Space) CommandList(item string) []*Command {
	if item == "" {
		return space.Entries
	}

	var result []*Command
	for _, command := range space.Entries {
		// match by label
		if command.Label == item {
			result = append(result, command)
			continue
		}

		// match by tag
		if command.Tagged(item) {
			result = append(result, command)
		}
	}
	return result
}

func (space *Space) commandFindPositionByLabel(commandLabel string) (int, error) {
	if commandLabel == "" {
		return -1, fmt.Errorf("could not search by empty label")
	}
	for i, command := range space.Entries {
		if command.Label == commandLabel {
			return i, nil
		}
	}
	return -1, fmt.Errorf("command with label '%s' not found", commandLabel)
}

func (space *Space) CommandFind(label string) (*Command, error) {
	pos, err := space.commandFindPositionByLabel(label)
	if err != nil {
		return nil, fmt.Errorf("find command: %v", err)
	}
	return space.Entries[pos], nil
}

func (space *Space) CommandDelete(command *Command) {
	pos, err := space.commandFindPositionByLabel(command.Label)
	if err != nil {
		log.Fatalf("delete command: %v", err)
	}
	space.UpdatedAt = UnixTimeNow()
	space.Entries = append(space.Entries[:pos], space.Entries[pos+1:]...)
}

func (space *Space) TagsList(filterTag string) []string {
	commands := space.CommandList(filterTag)

	result := []string{}

	tags := make(map[string]struct{})
	var found struct{}
	for _, command := range commands {
		for _, tag := range command.Tags {
			if _, ok := tags[tag]; !ok {
				tags[tag] = found
				result = append(result, tag)
			}
		}
	}

	return result
}

func (space *Space) SearchCommands(tag string, criteria string) ([]*Command, error) {
	if criteria == "" {
		return nil, fmt.Errorf("could not search with empty criteria")
	}

	var results []*Command
	for _, command := range space.Entries {
		if tag != "" && !command.Tagged(tag) {
			continue
		}
		if command.Matches(criteria) {
			results = append(results, command)
		}
	}
	return results, nil
}
