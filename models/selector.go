package models

import (
	"fmt"
	"log"
	"regexp"
)

type Selector struct {
	Item  string
	Space string
}

func ParseSelector(str string) (*Selector, error) {
	selector := Selector{}

	validRegex, err := regexp.Compile("^([a-z0-9-]+)?(@[a-z0-9-]+)?$")
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

	itemRegex, err := regexp.Compile("^[a-z0-9-]+")
	if err != nil {
		log.Fatal("Could not compile item regexp", err)
	}

	if spaceRegex.MatchString(str) {
		selector.Space = spaceRegex.FindString(str)[1:]
	} else {
		selector.Space = "default"
	}

	if itemRegex.MatchString(str) {
		selector.Item = itemRegex.FindString(str)
	}

	return &selector, nil
}
