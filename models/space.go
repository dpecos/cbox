package models

import (
	"github.com/satori/go.uuid"
)

type Space struct {
	ID    uuid.UUID
	Name  string
	Title string
}
