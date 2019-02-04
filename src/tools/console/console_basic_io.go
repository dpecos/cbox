package console

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/dplabs/cbox/src/tools/tty"
	bitflag "github.com/mvpninjas/go-bitflag"
	survey "gopkg.in/AlecAivazis/survey.v1"
	surveyCore "gopkg.in/AlecAivazis/survey.v1/core"
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

func init() {
	surveyCore.SetFancyIcons()
}

func ReadString(label string, opts ...bitflag.Flag) string {
	return readString(label, false, opts...)
}

func readString(label string, dieOnAbort bool, opts ...bitflag.Flag) string {
	var flags bitflag.Flag
	flags.Set(opts...)

	value, aborted := readStringDetails(label, "", flags, dieOnAbort)

	if aborted && flags.Isset(NOT_EMPTY_VALUES) {
		PrintError(MSG_EMPTY_NOT_ALLOWED)
		value = readString(label, true, opts...)
	} else if flags.Isset(NOT_EMPTY_VALUES) && strings.TrimSpace(value) == "" {
		PrintError(MSG_EMPTY_NOT_ALLOWED)
		value = readString(label, false, opts...)
	} else if flags.Isset(ONLY_VALID_CHARS) && !CheckValidChars(value) {
		PrintError(MSG_NOT_VALID_CHARS)
		value = readString(label, false, opts...)
	}

	return value
}

func EditString(label string, previousValue string, opts ...bitflag.Flag) string {
	value, aborted := editString(label, previousValue, false, opts...)
	return resolveEditionValue(previousValue, value, aborted)
}

func editString(label string, previousValue string, dieOnAbort bool, opts ...bitflag.Flag) (string, bool) {
	var flags bitflag.Flag
	flags.Set(opts...)

	value, aborted := readStringDetails(label, previousValue, flags, dieOnAbort)

	if aborted && flags.Isset(NOT_EMPTY_VALUES) {
		PrintError(MSG_EMPTY_NOT_ALLOWED)
		return editString(label, previousValue, true, opts...)
	} else if flags.Isset(ONLY_VALID_CHARS) && !CheckValidChars(value) {
		PrintError(MSG_NOT_VALID_CHARS)
		return editString(label, previousValue, false, opts...)
	}

	return value, aborted
}

func readStringDetails(label string, previousValue string, flags bitflag.Flag, dieOnAbort bool) (string, bool) {
	value := ""
	aborted := false

	if dieOnAbort {
		PrintWarning("cbox will terminate if you press Ctrl+C once more")
	}

	var prompt survey.Prompt

	help := ""
	if !flags.Isset(NOT_EMPTY_VALUES) {
		help = "Blank entry keeps previous value. Ctrl+C clears it"
	}

	if flags.Isset(MULTILINE) {
		prompt = &survey.Multiline{
			Message: label,
			Help:    help,
		}
	} else {
		prompt = &survey.Input{
			Message: label,
			Help:    help,
		}
	}

	if err := survey.AskOne(prompt, &value, nil); err != nil {
		if dieOnAbort {
			log.Fatal("cbox killed")
		}
		aborted = true
		value = ""
	}

	return value, aborted
}

func CheckValidChars(str string) bool {
	validCharsRegexp, err := regexp.Compile("^[a-z0-9-]*$")
	if err != nil {
		log.Fatalf("valid chars: could not compile regexp: %v", err)
	}

	return validCharsRegexp.MatchString(str)
}

func resolveEditionValue(previousValue string, newValue string, aborted bool) string {
	newValue = strings.TrimSpace(newValue)
	if aborted {
		// user wants to clear current value
		PrintInfo("Value cleared")
		return ""
	} else {
		if newValue == "" {
			PrintInfo(fmt.Sprintf("Using previous value: '%s'", previousValue))
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

	fmt.Println()

	return response
}

func PrintError(msg string) {
	fmt.Printf("%s\n", tty.ColorRed(msg))
}

func PrintSuccess(msg string) {
	fmt.Printf("%s\n", tty.ColorGreen(msg))
}

func PrintInfo(msg string) {
	fmt.Printf("%s\n", tty.ColorCyan(msg))
}

func PrintWarning(msg string) {
	fmt.Printf("%s %s\n", tty.ColorBgRed(" WARNING "), tty.ColorMagenta(msg))
}

func PrintAction(msg string) {
	fmt.Printf("** %s **\n\n", tty.ColorBoldYellow(msg))
}
