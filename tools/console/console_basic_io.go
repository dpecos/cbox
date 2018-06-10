package console

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
)

const MSG_EMPTY_TO_FINISH = "Empty line to finish"
const MSG_EDIT = "Ctrl+D to clear, Empty line to maintain"

func formatLabel(label string) string {
	return fmt.Sprintf("  %s: ", aurora.Blue(aurora.Bold(label)))
}

func formatLabelDetails(label string, details string) string {
	return fmt.Sprintf("  %s %s: ", aurora.Blue(aurora.Bold(label)), aurora.Blue("("+details+")"))
}

func formatPreviousValue(label string, value string) string {
	return fmt.Sprintf("  %s: %s", aurora.Cyan(label), value)
}

func formatQuestion(question string, options string) string {
	return fmt.Sprintf("%s %s: ", aurora.Magenta(aurora.Bold(question)), aurora.Magenta(options))
}

func ReadString(label string) string {
	fmt.Printf(formatLabel(label))
	value, _ := readString()
	return value
}

func ReadStringMulti(label string) string {
	val, _ := readStringMulti(label, MSG_EMPTY_TO_FINISH)
	return val
}

func readString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	val, err := reader.ReadString('\n')
	return strings.TrimSpace(val), err
}

func readStringDetails(label string, details string) (string, bool) {
	fmt.Printf(formatLabelDetails(label, details))
	val, err := readString()
	if err == io.EOF {
		return "", true
	}
	return val, false
}

func readStringMulti(label string, details string) (string, bool) {

	if details != "" {
		fmt.Println(formatLabelDetails(label, details))
	} else {
		fmt.Println(formatLabel(label))
	}

	reader := bufio.NewReader(os.Stdin)

	arr := make([]string, 0)
	for {
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			return "", true
		}
		if text != "\n" {
			arr = append(arr, text)
			label = ""
		} else {
			break
		}
	}

	val := strings.Join(arr, "")
	return strings.TrimSpace(val), false
}

func EditString(label string, previousValue string) string {
	fmt.Println(formatPreviousValue("Previous value of "+strings.ToLower(label), previousValue))
	value, aborted := readStringDetails(label, MSG_EDIT)
	fmt.Println()

	return resolveEditionValue(previousValue, value, aborted)
}

func EditStringMulti(label string, previousValue string) string {
	fmt.Println(formatPreviousValue("Previous value of "+strings.ToLower(label), previousValue))
	val, aborted := readStringMulti("New value for "+label, MSG_EDIT)
	fmt.Println()

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
	fmt.Printf(formatQuestion(label, "y/n"))
	val, _ := readString()
	fmt.Println()
	return val == "y"
}

func PrintError(msg string) {
	fmt.Printf("%s\n\n", aurora.Red(msg))
}

func PrintSuccess(msg string) {
	fmt.Printf("%s\n\n", aurora.Green(msg))
}
