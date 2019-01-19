package models

type Meta struct {
	ID        string    `json:"id"`
	Selector  *Selector `json:"-"`
	UpdatedAt UnixTime  `json:"updated-at" dynamodbav:",unixtime"`
	CreatedAt UnixTime  `json:"created-at" dynamodbav:",unixtime"`
}
