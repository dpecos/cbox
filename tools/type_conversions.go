package tools

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func StringToInt(str string) int64 {
	param, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return param
}

func DateToString(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
