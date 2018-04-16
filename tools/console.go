package tools

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/peterh/liner"
)

func ReadStringMulti(label string) string {
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)

	if label != "" {
		label += " (empty line to finish): "
	}

	arr := make([]string, 0)
	for {
		text, err := line.Prompt(label)
		if err != nil {
			if err == liner.ErrPromptAborted {
				line.Close()
				os.Exit(0)
			} else {
				log.Fatal(err)
			}
		}
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
	fmt.Printf("%s: ", label)

	reader := bufio.NewReader(os.Stdin)
	val, _ := reader.ReadString('\n')
	return strings.TrimSpace(val)
}
