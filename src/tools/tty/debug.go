package tty

import (
	"fmt"
)

func Debug(msg string) {
	fmt.Printf("%s\n", ColorBoldBlack(msg))
}
