package tty

import "fmt"

const (
	COLOR_BLACK   = 30
	COLOR_RED     = 31
	COLOR_GREEN   = 32
	COLOR_YELLOW  = 33
	COLOR_BLUE    = 34
	COLOR_MAGENTA = 35
	COLOR_CYAN    = 36
	COLOR_WHITE   = 37
	COLOR_RESET   = 0
)

var DisableColors = false

func colorize(color int, bold bool, bg bool, str string) string {
	if DisableColors {
		return str
	}

	if bg {
		color += 10
	}
	if bold {
		return fmt.Sprintf("\x1b[%d;1m%s\x1b[%dm", color, str, COLOR_RESET)
	} else {
		return fmt.Sprintf("\x1b[%dm%s\x1b[%dm", color, str, COLOR_RESET)
	}
}

func ColorBlack(str string) string {
	return colorize(COLOR_BLACK, false, false, str)
}
func ColorRed(str string) string {
	return colorize(COLOR_RED, false, false, str)
}
func ColorGreen(str string) string {
	return colorize(COLOR_GREEN, false, false, str)
}
func ColorYellow(str string) string {
	return colorize(COLOR_YELLOW, false, false, str)
}
func ColorBlue(str string) string {
	return colorize(COLOR_BLUE, false, false, str)
}
func ColorMagenta(str string) string {
	return colorize(COLOR_MAGENTA, false, false, str)
}
func ColorCyan(str string) string {
	return colorize(COLOR_CYAN, false, false, str)
}
func ColorWhite(str string) string {
	return colorize(COLOR_WHITE, false, false, str)
}

func ColorBoldBlack(str string) string {
	return colorize(COLOR_BLACK, true, false, str)
}
func ColorBoldRed(str string) string {
	return colorize(COLOR_RED, true, false, str)
}
func ColorBoldGreen(str string) string {
	return colorize(COLOR_GREEN, true, false, str)
}
func ColorBoldYellow(str string) string {
	return colorize(COLOR_YELLOW, true, false, str)
}
func ColorBoldBlue(str string) string {
	return colorize(COLOR_BLUE, true, false, str)
}
func ColorBoldMagenta(str string) string {
	return colorize(COLOR_MAGENTA, true, false, str)
}
func ColorBoldCyan(str string) string {
	return colorize(COLOR_CYAN, true, false, str)
}
func ColorBoldWhite(str string) string {
	return colorize(COLOR_WHITE, true, false, str)
}

func ColorBgBlack(str string) string {
	return colorize(COLOR_BLACK, false, true, str)
}
func ColorBgRed(str string) string {
	return colorize(COLOR_RED, false, true, str)
}
func ColorBgGreen(str string) string {
	return colorize(COLOR_GREEN, false, true, str)
}
func ColorBgYellow(str string) string {
	return colorize(COLOR_YELLOW, false, true, str)
}
func ColorBgBlue(str string) string {
	return colorize(COLOR_BLUE, false, true, str)
}
func ColorBgMagenta(str string) string {
	return colorize(COLOR_MAGENTA, false, true, str)
}
func ColorBgCyan(str string) string {
	return colorize(COLOR_CYAN, false, true, str)
}
func ColorBgWhite(str string) string {
	return colorize(COLOR_WHITE, false, true, str)
}

func ColorBgBoldBlack(str string) string {
	return colorize(COLOR_BLACK, true, true, str)
}
func ColorBgBoldRed(str string) string {
	return colorize(COLOR_RED, true, true, str)
}
func ColorBgBoldGreen(str string) string {
	return colorize(COLOR_GREEN, true, true, str)
}
func ColorBgBoldYellow(str string) string {
	return colorize(COLOR_YELLOW, true, true, str)
}
func ColorBgBoldBlue(str string) string {
	return colorize(COLOR_BLUE, true, true, str)
}
func ColorBgBoldMagenta(str string) string {
	return colorize(COLOR_MAGENTA, true, true, str)
}
func ColorBgBoldCyan(str string) string {
	return colorize(COLOR_CYAN, true, true, str)
}
func ColorBgBoldWhite(str string) string {
	return colorize(COLOR_WHITE, true, true, str)
}
