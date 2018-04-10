package tools

import (
	"log"
	"strconv"
)

func StringToInt(str string) int64 {
	param, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return param
}
