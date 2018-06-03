package models

import (
	"fmt"
	"log"
	"regexp"
)

type Selector struct {
	ID    string
	Tag   string
	Space string
}

func ParseSelector(str string) (*Selector, error) {
	selector := Selector{}

	validRegex, err := regexp.Compile("^(#[a-z0-9-]+)?([a-z0-9-]+)?(@[a-z0-9-]+)?$")
	if err != nil {
		log.Fatal("Could not compile space regexp", err)
	}

	if !validRegex.MatchString(str) {
		return nil, fmt.Errorf("Invalid selector %s", str)
	}

	spaceRegex, err := regexp.Compile("@[a-z0-9-]+$")
	if err != nil {
		log.Fatal("Could not compile space regexp", err)
	}

	tagRegex, err := regexp.Compile("^#[a-z0-9-]+")
	if err != nil {
		log.Fatal("Could not compile tag regexp", err)
	}

	idRegex, err := regexp.Compile("^[a-z0-9-]+")
	if err != nil {
		log.Fatal("Could not compile id regexp", err)
	}

	if spaceRegex.MatchString(str) {
		selector.Space = spaceRegex.FindString(str)[1:]
	} else {
		selector.Space = "default"
	}

	if tagRegex.MatchString(str) {
		selector.Tag = tagRegex.FindString(str)[1:]
	}

	if idRegex.MatchString(str) {
		selector.ID = idRegex.FindString(str)
	}

	return &selector, nil
}
