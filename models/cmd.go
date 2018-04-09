package models

import (
	"time"
)

type Cmd struct {
	ID          int
	Cmd         string
	Title       string
	Description string
	URL         string
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
