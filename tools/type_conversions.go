package tools

import (
	"log"
	"strconv"

	"github.com/gofrs/uuid"
)

func StringToInt(str string) int64 {
	number, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatalf("str-int: error converting: %v", err)
	}
	return number
}

func StringToUUID(str string) uuid.UUID {
	id, err := uuid.FromString(str)
	if err != nil {
		log.Fatalf("str-uuid: could not parse UUID: %v", err)
	}
	return id
}
