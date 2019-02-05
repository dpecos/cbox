package tty

import (
	"fmt"
)

var DisableOutput = false

func Print(format string, args ...interface{}) {
	if !DisableOutput {
		if len(args) != 0 {
			fmt.Printf(format, args...)
		} else {
			fmt.Print(format)
		}
	}
}

func Debug(msg string) {
	Print("%s\n", ColorBoldBlack(msg))
}
