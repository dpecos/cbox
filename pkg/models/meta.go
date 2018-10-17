package models

import (
	"github.com/gofrs/uuid"
)

type Meta struct {
	ID        uuid.UUID `json:"id" dynamodbav:"-"`
	UpdatedAt UnixTime  `json:"updated-at" dynamodbav:",unixtime"`
	CreatedAt UnixTime  `json:"created-at" dynamodbav:",unixtime"`
}
