package console

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	bitflag "github.com/mvpninjas/go-bitflag"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

const MSG_EDIT = "Ctrl+D to clear, Empty line to maintain"
const MSG_EMPTY_NOT_ALLOWED = "Empty value not allowed, please try again"
const MSG_NOT_VALID_CHARS = "Value contains not valid chars, please try again"

type Flag byte

const (
	NOT_EMPTY_VALUES bitflag.Flag = 1 << bitflag.Flag(iota)
	MULTILINE
	ONLY_VALID_CHARS
)

func ReadString(label string, opts ...bitflag.Flag) string {
	var flags bitflag.Flag
	flags.Set(opts...)

	value, aborted := readStringDetails(label, "", flags)

	if aborted && flags.Isset(NOT_EMPTY_VALUES) {
		PrintError(MSG_EMPTY_NOT_ALLOWED)
		value = ReadString(label, opts...)
	} else if flags.Isset(NOT_EMPTY_VALUES) && strings.TrimSpace(value) == "" {
		PrintError(MSG_EMPTY_NOT_ALLOWED)
		value = ReadString(label, opts...)
	} else if flags.Isset(ONLY_VALID_CHARS) && !checkValidChars(value) {
		PrintError(MSG_NOT_VALID_CHARS)
		value = ReadString(label, opts...)
	}

	return resolveEditionValue("", value, aborted)
}

func EditString(label string, previousValue string, opts ...bitflag.Flag) string {
	var flags bitflag.Flag
	flags.Set(opts...)

	value, aborted := readStringDetails(label, previousValue, flags)

	if aborted && flags.Isset(NOT_EMPTY_VALUES) {
		PrintError(MSG_EMPTY_NOT_ALLOWED)
		value = EditString(label, previousValue, opts...)
	} else if flags.Isset(NOT_EMPTY_VALUES) && strings.TrimSpace(value) == "" {
		PrintError(MSG_EMPTY_NOT_ALLOWED)
		value = EditString(label, previousValue, opts...)
	} else if flags.Isset(ONLY_VALID_CHARS) && !checkValidChars(value) {
		PrintError(MSG_NOT_VALID_CHARS)
		value = EditString(label, previousValue, opts...)
	}

	return resolveEditionValue(previousValue, value, aborted)
}

func readStringDetails(label string, previousValue string, flags bitflag.Flag) (string, bool) {
	value := ""
	aborted := false

	var prompt survey.Prompt

	if flags.Isset(MULTILINE) {
		prompt = &survey.MultilineInput{
			Message: label,
			Default: previousValue,
		}
		// prompt = &survey.Editor{
		// 	Message:       label,
		// 	Default:       previousValue,
		// 	AppendDefault: previousValue != "",
		// }
	} else {
		prompt = &survey.Input{
			Message: label,
			Default: previousValue,
		}
	}

	if err := survey.AskOne(prompt, &value, nil); err != nil {
		aborted = true
	}

	return value, aborted
}

func checkValidChars(str string) bool {
	validCharsRegexp, err := regexp.Compile("^[a-z0-9-]*$")
	if err != nil {
		log.Fatalf("valid chars: could not compile regexp: %v", err)
	}

	if !validCharsRegexp.MatchString(str) {
		return false
	}

	return true
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
	response := false

	prompt := &survey.Confirm{
		Message: label,
	}
	survey.AskOne(prompt, &response, nil)

	return response
}

func PrintError(msg string) {
	fmt.Printf("%s\n", ColorRed(msg))
}

func PrintSuccess(msg string) {
	fmt.Printf("%s\n", ColorGreen(msg))
}

func PrintInfo(msg string) {
	fmt.Printf("%s\n", ColorCyan(msg))
}

func PrintWarning(msg string) {
	fmt.Printf("%s %s\n", ColorBgRed(" WARNING "), ColorMagenta(msg))
}

func Debug(msg string) {
	fmt.Printf("%s\n", ColorBoldBlack(msg))
}

func PrintAction(msg string) {
	fmt.Printf("%s\n\n", msg)
}
