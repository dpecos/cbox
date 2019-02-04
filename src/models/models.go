package models

import (
	"net/http"
	"net/url"
	"time"
)

type UnixTime time.Time

const (
	TypeNone = iota
	TypeUser
	TypeOrganization
)

type Selector struct {
	Item          string
	NamespaceType int
	Namespace     string
	Space         string
}

type CBox struct {
	Spaces  []*Space
	Version string
	Build   string
}

type Meta struct {
	ID        string    `json:"id"`
	Selector  *Selector `json:"-"`
	UpdatedAt UnixTime  `json:"updated-at" dynamodbav:",unixtime"`
	CreatedAt UnixTime  `json:"created-at" dynamodbav:",unixtime"`
}

type Space struct {
	Meta
	Label       string     `json:"label"`
	Description string     `json:"description"`
	Entries     []*Command `json:"entries" dynamodbav:"-"`
}

type Command struct {
	Meta
	Label       string   `json:"label"`
	Code        string   `json:"code"`
	Description string   `json:"description"`
	URL         string   `json:"url" dynamodbav:",omitempty"`
	Tags        []string `json:"tags" dynamodbav:",omitempty"`
}

type Cloud struct {
	Environment string
	ServerKey   string
	UserID      string
	Login       string
	Name        string
	Token       string
	URL         string
	BaseURL     *url.URL
	HttpClient  *http.Client
	Cbox        *CBox
}
