package models

import (
	"fmt"
	"log"
	"regexp"
)

type Selector struct {
	Item         string
	Organization string
	User         string
	Space        string
}

func (selector *Selector) String() string {
	if selector == nil {
		return ""
	}
	if selector.User != "" {
		return fmt.Sprintf("%s@%s:%s", selector.Item, selector.User, selector.Space)
	}
	return fmt.Sprintf("%s@%s", selector.Item, selector.Space)
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

func check(selector *Selector, str string, item bool, user bool, space bool) error {
	if item && selector.Item == "" {
		return fmt.Errorf("item not specified in the selector: '%s'", str)
	}
	if user && selector.User == "" {
		return fmt.Errorf("user not specified in the selector: '%s'", str)
	}
	if space && selector.Space == "" {
		return fmt.Errorf("space not specified in the selector: '%s'", str)
	}
	return nil
}

func ParseSelectorMandatorySpace(str string) (*Selector, error) {
	selector, err := parseSelector(str)

	if err == nil {
		if err := check(selector, str, false, false, true); err != nil {
			return nil, err
		}
	}

	return selector, err
}

func ParseSelectorMandatoryItem(str string) (*Selector, error) {
	selector, err := ParseSelector(str)

	if err == nil {
		if err := check(selector, str, true, false, false); err != nil {
			return nil, err
		}
	}

	return selector, err
}

func ParseSelectorForCloudCommand(str string) (*Selector, error) {
	selector, err := ParseSelector(str)

	if err == nil {
		if err := check(selector, str, true, true, true); err != nil {
			return nil, err
		}
	}

	return selector, err
}

func parseSelector(str string) (*Selector, error) {

	selectorRegexp, err := regexp.Compile("^(?P<item>[a-z0-9-]+)?(@(?P<organization>[a-z0-9-]+/)?((?P<user>[a-z0-9-]+):)?(?P<space>[a-z0-9-]+))?$")
	if err != nil {
		log.Fatalf("parse selector: could not compile selector regexp: %v", err)
	}

	if !selectorRegexp.MatchString(str) {
		return nil, fmt.Errorf("parse selector: invalid selector: '%s'", str)
	}

	match := selectorRegexp.FindStringSubmatch(str)

	selectorMap := make(map[string]string)
	for i, name := range selectorRegexp.SubexpNames() {
		if i > 0 && i <= len(match) {
			selectorMap[name] = match[i]
		}
	}

	selector := Selector{
		Item:         selectorMap["item"],
		Organization: selectorMap["organization"],
		User:         selectorMap["user"],
		Space:        selectorMap["space"],
	}

	return &selector, nil
}
