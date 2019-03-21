package tty

import (
	"fmt"
	"log"

	survey "gopkg.in/AlecAivazis/survey.v1"
	surveyCore "gopkg.in/AlecAivazis/survey.v1/core"
)

var (
	DisableOutput = false

	MockTTY      = false
	MockedOutput = ""
	MockedInput  = []string{}
)

func init() {
	surveyCore.SetFancyIcons()
}

func Print(format string, args ...interface{}) {
	if !DisableOutput {
		var nl string
		if len(args) != 0 {
			nl = fmt.Sprintf(format, args...)
		} else {
			nl = fmt.Sprintf(format)
		}
		if MockTTY {
			MockedOutput = MockedOutput + nl
		} else {
			fmt.Print(nl)
		}
	}
}

func Read(label string, help string, multiline bool) (string, error) {
	if MockTTY {
		if len(MockedInput) > 0 {
			value := MockedInput[0]
			MockedInput = MockedInput[1:]
			return value, nil
		}
		log.Fatalf("input mocked but not enough values provided for input - Mocked output until this moment: \n %s", MockedOutput)
	}

	var prompt survey.Prompt

	if multiline {
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

	var value string
	err := survey.AskOne(prompt, &value, nil)

	return value, err
}

func Confirm(label string) bool {
	if MockTTY {
		if len(MockedInput) > 0 {
			value := MockedInput[0]
			MockedInput = MockedInput[1:]
			return value == "y"
		}
		log.Fatalf("input mocked but not enough values provided for input")
	}

	response := false

	prompt := &survey.Confirm{
		Message: label,
	}
	survey.AskOne(prompt, &response, nil)

	Print("\n")

	return response
}

func Debug(msg string) {
	Print("%s\n", ColorBoldBlack(msg))
}
