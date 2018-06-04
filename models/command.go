package models

import (
	"time"
)

type Command struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Tags        []string  `json:"tags"`
	UpdatedAt   time.Time `json:"updated-at"`
	CreatedAt   time.Time `json:"created-at"`
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
