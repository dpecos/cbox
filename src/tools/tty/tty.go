package tty

import (
	"fmt"
)

var (
	DisableOutput = false

	MockTTY      = false
	MockedOutput = ""
)

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

func Debug(msg string) {
	Print("%s\n", ColorBoldBlack(msg))
}
