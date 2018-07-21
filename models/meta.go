package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Meta struct {
	ID        uuid.UUID `json:"id" dynamodbav:"-"`
	UpdatedAt time.Time `json:"updated-at" dynamodbav:",unixtime"`
	CreatedAt time.Time `json:"created-at" dynamodbav:",unixtime"`
}
