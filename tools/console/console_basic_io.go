package console

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
	bitflag "github.com/mvpninjas/go-bitflag"
)

const MSG_EMPTY_TO_FINISH = "Empty line to finish"
const MSG_EDIT = "Ctrl+D to clear, Empty line to maintain"
const MSG_EMPTY_NOT_ALLOWED = "Empty value not allowed, please try again"

type Flag byte

const (
	NOT_EMPTY_VALUES bitflag.Flag = 1 << bitflag.Flag(iota)
	MULTILINE
)

func formatLabel(label string) string {
	return fmt.Sprintf("%s: ", aurora.Blue(aurora.Bold(label)))
}

func formatLabelDetails(label string, details string) string {
	return fmt.Sprintf("%s %s: ", aurora.Blue(aurora.Bold(label)), aurora.Blue("("+details+")"))
}

func formatPreviousValue(label string, value string) string {
	return fmt.Sprintf("%s: %s", aurora.Cyan(label), value)
}

func formatQuestion(question string, options string) string {
	return fmt.Sprintf("%s %s: ", aurora.Magenta(aurora.Bold(question)), aurora.Magenta(options))
}

func ReadString(label string, opts ...bitflag.Flag) string {
	return readStringDetails(label, "", opts...)
}

func ReadStringDetails(label string, details string, opts ...bitflag.Flag) string {
	return readStringDetails(label, details, opts...)
}

func readStringDetails(label string, details string, opts ...bitflag.Flag) string {
	var flags bitflag.Flag
	flags.Set(opts...)

	if flags.Isset(MULTILINE) && details == "" {
		details = MSG_EMPTY_TO_FINISH
	}

	value, _ := readString(label, details, flags.Isset(MULTILINE))

	if flags.Isset(NOT_EMPTY_VALUES) && strings.TrimSpace(value) == "" {
		PrintError(MSG_EMPTY_NOT_ALLOWED)
		value, _ = readString(label, details, flags.Isset(MULTILINE))
	}

	return value
}

func readString(label string, details string, multiline bool) (string, bool) {
	if details != "" {
		fmt.Print(formatLabelDetails(label, details))
	} else {
		fmt.Print(formatLabel(label))
	}

	if multiline {
		fmt.Println()
	}

	if multiline {
		return readStringMulti()
	} else {
		return readStringSimple()
	}
}

func readStringSimple() (string, bool) {
	reader := bufio.NewReader(os.Stdin)
	val, err := reader.ReadString('\n')

	if err == io.EOF {
		return "", true
	}
	return strings.TrimSpace(val), false
}

func readStringMulti() (string, bool) {

	reader := bufio.NewReader(os.Stdin)

	arr := make([]string, 0)
	for {
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			return "", true
		}
		if text != "\n" {
			arr = append(arr, text)
		} else {
			break
		}
	}

	val := strings.Join(arr, "")
	return strings.TrimSpace(val), false
}

func EditString(label string, previousValue string, opts ...bitflag.Flag) string {
	var flags bitflag.Flag
	flags.Set(opts...)

	value := editString(label, previousValue, flags.Isset(MULTILINE))

	if flags.Isset(NOT_EMPTY_VALUES) && strings.TrimSpace(value) == "" {
		PrintError(MSG_EMPTY_NOT_ALLOWED)
		value = editString(label, previousValue, flags.Isset(MULTILINE))
	}

	return value
}

func editString(label string, previousValue string, multiline bool) string {
	fmt.Println(formatPreviousValue("Previous value of "+label, previousValue))
	val, aborted := readString("New value for "+label, MSG_EDIT, multiline)
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
	val, _ := readStringSimple()
	fmt.Println()
	return val == "y"
}

func PrintError(msg string) {
	fmt.Printf("%s\n\n", aurora.Red(msg))
}

func PrintSuccess(msg string) {
	fmt.Printf("%s\n\n", aurora.Green(msg))
}
