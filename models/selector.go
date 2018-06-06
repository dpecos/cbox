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

const DEFAULT_SPACE_NAME = "default"

func ParseSelector(str string) (*Selector, error) {
	selector, err := parseSelector(str)

	if err == nil {
		if selector.Space == "" {
			selector.Space = DEFAULT_SPACE_NAME
		}
	}

	return selector, err
}

func ParseSelectorMandatorySpace(str string) (*Selector, error) {
	selector, err := parseSelector(str)

	if err == nil {
		if selector.Space == "" {
			return nil, fmt.Errorf("Space not specified in the selector '%s'", str)
		}
	}

	return selector, nil
}

func parseSelector(str string) (*Selector, error) {
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
	}

	if itemRegex.MatchString(str) {
		selector.Item = itemRegex.FindString(str)
	}

	return &selector, nil
}
