package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/peterh/liner"
)

func ReadStringMulti(label string) string {
	val, _ := readStringMulti(label, false)
	return val
}

func readStringMulti(label string, allowAbortEdition bool) (string, bool) {
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)

	if label != "" {
		if allowAbortEdition {
			label += " (Ctrl+C to clear current value): "
		} else {
			label += " (empty line to finish): "
		}
	}

	arr := make([]string, 0)
	for {
		text, err := line.Prompt(label)
		aborted := handleLinerError(line, err, allowAbortEdition)

		if aborted {
			return "", true
		}

		if len(text) != 0 {
			arr = append(arr, text)
			label = ""
		} else {
			break
		}
	}

	val := strings.Join(arr, "\n")
	return strings.TrimSpace(val), false
}

func handleLinerError(line *liner.State, err error, allowAbortEdition bool) bool {
	if err != nil {
		if err == liner.ErrPromptAborted {
			if !allowAbortEdition {
				line.Close()
				os.Exit(0)
			} else {
				return true
			}
		} else {
			log.Fatal(err)
		}
	}
	return false
}

func ReadString(label string) string {
	fmt.Printf("%s: ", label)

	reader := bufio.NewReader(os.Stdin)
	val, _ := reader.ReadString('\n')
	return strings.TrimSpace(val)
}

func EditString(label string, previousValue string) string {
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)

	label += " (Ctrl+C to clear current value)"

	val, err := line.PromptWithSuggestion(fmt.Sprintf("%s: ", label), previousValue, len(previousValue))
	aborted := handleLinerError(line, err, true)

	return resolveEditionValue(previousValue, val, aborted)
}

func EditStringMulti(label string, previousValue string) string {
	label = strings.ToLower(label)
	fmt.Printf("Previous value for %s:\n%s\n", label, previousValue)
	val, aborted := readStringMulti("New value for "+label, true)

	return resolveEditionValue(previousValue, val, aborted)
}

func resolveEditionValue(previousValue string, newValue string, aborted bool) string {
	newValue = strings.TrimSpace(newValue)
	if aborted {
		// user wants to clear current value
		return ""
	} else {
		if newValue == "" {
			// user wants to keep current value
			return previousValue
		} else {
			return newValue
		}
	}
}

func Confirm(label string) bool {
	return ReadString(label+" (y/n)") == "y"
}

func PrintError(msg string) {
	fmt.Println(aurora.Red(msg))
}
