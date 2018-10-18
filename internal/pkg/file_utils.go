package pkg

import (
	"log"
	"os"
)

func CreateDirectoryIfNotExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			log.Fatalf("could not create directory: %v", err)
		}
		return true
	}
	return false
}

func CreateFileIfNotExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0600)
		if err != nil {
			log.Fatalf("could not create file: %v", err)
		}
		return true
	}
	return false
}
