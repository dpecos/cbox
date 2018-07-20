package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dpecos/cbox/tools/console"
	uuid "github.com/satori/go.uuid"
)

type Space struct {
	ID          uuid.UUID `json:"id"`
	Label       string    `json:"label"`
	Description string    `json:"description"`
	Entries     []Command `json:"entries,omitempty" dynamodbav:",omitempty"`
	UpdatedAt   time.Time `json:"updated-at"`
}

func commandPresentInSapce(space *Space, commandLabel string) bool {
	for _, cmd := range space.Entries {
		if commandLabel == cmd.Label {
			return true
		}
	}
	return false
}

func (space *Space) CommandAdd(command *Command) {
	for commandPresentInSapce(space, command.Label) {
		console.PrintError("Label already found in space. Try a different one")
		command.Label = strings.ToLower(console.ReadString("Label"))
	}
	now := time.Now()
	command.CreatedAt = now
	command.UpdatedAt = now
	space.Entries = append(space.Entries, *command)
}

func (space *Space) CommandEdit(command *Command, previousLabel string) {
	if command.Label != previousLabel {
		newLabel := command.Label
		command.Label = previousLabel
		for commandPresentInSapce(space, newLabel) {
			console.PrintError("Label already found in space. Try a different one")
			newLabel = strings.ToLower(console.ReadString("Label"))
		}
		command.Label = newLabel
	}
	command.UpdatedAt = time.Now()
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

func (space *Space) commandFindPositionByID(commandID uuid.UUID) (int, error) {
	if commandID == uuid.Nil {
		return -1, fmt.Errorf("could not search by empty ID")
	}
	for i, command := range space.Entries {
		if command.ID == commandID {
			return i, nil
		}
	}
	return -1, fmt.Errorf("command with ID '%s' not found", commandID)
}

func (space *Space) CommandFind(commandLocator string) *Command {
	pos, err := space.commandFindPositionByLabel(commandLocator)
	if err != nil {
		id := uuid.FromStringOrNil(commandLocator)
		pos, err = space.commandFindPositionByID(id)
		if err != nil {
			log.Fatalf("find command: %v", err)
		}
	}
	return &space.Entries[pos]
}

func (space *Space) CommandDelete(command *Command) {
	pos, err := space.commandFindPositionByID(command.ID)
	if err != nil {
		log.Fatalf("delete command: %v", err)
	}

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
