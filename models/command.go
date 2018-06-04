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
