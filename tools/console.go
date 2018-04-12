package tools

import (
	"log"
	"strings"

	"github.com/peterh/liner"
)

func ReadStringMulti(label string) string {
	arr := make([]string, 0)

	for {
		text := ReadString(label + " (empty line to finish)")
		if len(text) != 0 {
			arr = append(arr, text)
			label = ""
		} else {
			break
		}
	}

	val := strings.Join(arr, "\n")
	return strings.TrimSpace(val)
}

func ReadString(label string) string {
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)

	if label != "" {
		label += ": "
	}

	val, err := line.Prompt(label)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(val)
}
