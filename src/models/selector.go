package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/spf13/viper"
)

type Selector struct {
	Item      string
	Namespace string
	Space     string
}

func (selector *Selector) String() string {
	if selector == nil {
		return ""
	}
	return fmt.Sprintf("%s@%s%s", selector.Item, selector.Namespace, selector.Space)

	return fmt.Sprintf("%s@%s", selector.Item, selector.Space)
}

func ParseSelector(str string) (*Selector, error) {
	selector, err := parseSelector(str)

	if err == nil {
		if selector.Space == "" {
			selector.Space = viper.GetString("cbox.default-space")
		}
	}

	return selector, err
}

func check(selector *Selector, str string, item bool, namespace bool, space bool) error {
	if item && selector.Item == "" {
		return fmt.Errorf("item not specified in the selector: '%s'", str)
	}
	if namespace && selector.Namespace == "" {
		return fmt.Errorf("namespace not specified in the selector: '%s'", str)
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

func ParseSelectorForCloud(str string) (*Selector, error) {
	selector, err := ParseSelector(str)

	if err == nil {
		if err := check(selector, str, false, true, true); err != nil {
			return nil, err
		}
	}

	return selector, err
}

func parseSelector(str string) (*Selector, error) {

	// selectorRegexp, err := regexp.Compile("^(?P<item>[a-z0-9-]+)?(@((?P<namespace>[a-z0-9-]+)(?P<qualifier>[:/]))?(?P<space>[a-z0-9-]+))?$")
	selectorRegexp, err := regexp.Compile("^(?P<item>[a-z0-9-]+)?(@(?P<namespace>[a-z0-9-]+[:/])?(?P<space>[a-z0-9-]+))?$")
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
		Item:      selectorMap["item"],
		Namespace: selectorMap["namespace"],
		Space:     selectorMap["space"],
	}

	// if qualifier, ok := selectorMap["qualifier"]; ok {
	// 	if qualifier == ":" {
	// 		selector.Qualifier = "USER"
	// 	} else if qualifier == "/" {
	// 		selector.Qualifier = "ORGANIZATION"
	// 	}
	// 	selector.Namespace = selectorMap["namespace"]
	// }

	return &selector, nil
}
