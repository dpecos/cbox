package tools

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/satori/go.uuid"
)

func StringToInt(str string) int64 {
	number, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func StringToUUID(str string) uuid.UUID {
	id, err := uuid.FromString(str)
	if err != nil {
		log.Fatal("Could not parse UUID", err)
	}
	return id
}

func DateToString(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
