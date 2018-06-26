package tools

import "os"

func CreateDirectoryIfNotExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
		return true
	}
	return false
}

func CreateFileIfNotExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0600)
		return true
	}
	return false
}
