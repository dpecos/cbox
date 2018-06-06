package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dpecos/cbox/tools/console"
)

type Space struct {
	Name    string    `json:"name"`
	Title   string    `json:"title"`
	Entries []Command `json:"entries"`
}

func commandPresentInSapce(space *Space, commandID string) bool {
	for _, cmd := range space.Entries {
		if commandID == cmd.ID {
			return true
		}
	}
	return false
}

func (space *Space) CommandAdd(command *Command) {
	for commandPresentInSapce(space, command.ID) {
		console.PrintError("ID already found in space. Try a different one")
		command.ID = strings.ToLower(console.ReadString("ID"))
	}
	command.CreatedAt = time.Now()
	command.UpdatedAt = time.Now()
	space.Entries = append(space.Entries, *command)
}

func (space *Space) CommandEdit(command *Command, previousID string) {
	if command.ID != previousID {
		newID := command.ID
		command.ID = previousID
		for commandPresentInSapce(space, newID) {
			console.PrintError("ID already found in space. Try a different one")
			newID = strings.ToLower(console.ReadString("ID"))
		}
		command.ID = newID
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

func (space *Space) commandFindPosition(commandId string) (int, error) {
	for i, command := range space.Entries {
		if command.ID == commandId {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Could not find command with ID %s", commandId)
}

func (space *Space) CommandFind(commandId string) *Command {
	pos, err := space.commandFindPosition(commandId)
	if err != nil {
		log.Fatal(err)
	}
	return &space.Entries[pos]
}

func (space *Space) CommandDelete(command *Command) {
	pos, err := space.commandFindPosition(command.ID)
	if err != nil {
		log.Fatal(err)
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
