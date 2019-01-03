package models

type Meta struct {
	UpdatedAt UnixTime `json:"updated-at" dynamodbav:",unixtime"`
	CreatedAt UnixTime `json:"created-at" dynamodbav:",unixtime"`
}
