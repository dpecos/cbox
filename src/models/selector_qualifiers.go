package models

import (
	"fmt"
)

func SOrganization(organization string) string {
	return fmt.Sprintf("%s/", organization)
}
func SUser(username string) string {
	return fmt.Sprintf("%s:", username)
}
func SNamespace(namespace string) string {
	if namespace == "" {
		return ""
	}
	return namespace[:len(namespace)-1]
}
