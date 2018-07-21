package models

import (
	"time"
)

type Command struct {
	Meta
	Label       string   `json:"label"`
	Code        string   `json:"code"`
	Description string   `json:"description"`
	Details     string   `json:"details" dynamodbav:",omitempty"`
	URL         string   `json:"url" dynamodbav:",omitempty"`
	Tags        []string `json:"tags" dynamodbav:",omitempty"`
}

func (command *Command) TagAdd(tag string) {
	found := false

	for _, t := range command.Tags {
		if t == tag {
			found = true
			break
		}
	}

	if !found {
		command.Tags = append(command.Tags, tag)
		command.UpdatedAt = time.Now()
	}
}

func (command *Command) TagDelete(tag string) {
	found := -1

	for i, t := range command.Tags {
		if t == tag {
			found = i
			break
		}
	}

	if found != -1 {
		command.Tags = append(command.Tags[:found], command.Tags[found+1:]...)
		command.UpdatedAt = time.Now()
	}
}
