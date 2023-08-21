package ansi

import (
	"fmt"
	"strconv"
)

const ESC = "\033"

// Text Styling
const (
	Reset      = ESC + "[0m"
	Bold       = ESC + "[1m"
	Faint      = ESC + "[2m"
	Italic     = ESC + "[3m"
	Underline  = ESC + "[4m"
	Blink      = ESC + "[5m"
	Reverse    = ESC + "[7m"
	CrossedOut = ESC + "[9m"
)

// Text Color
const (
	Black   = ESC + "[30m"
	Red     = ESC + "[31m"
	Green   = ESC + "[32m"
	Yellow  = ESC + "[33m"
	Blue    = ESC + "[34m"
	Magenta = ESC + "[35m"
	Cyan    = ESC + "[36m"
	White   = ESC + "[37m"

	BrightBlack   = ESC + "[90m"
	BrightRed     = ESC + "[91m"
	BrightGreen   = ESC + "[92m"
	BrightYellow  = ESC + "[93m"
	BrightBlue    = ESC + "[94m"
	BrightMagenta = ESC + "[95m"
	BrightCyan    = ESC + "[96m"
	BrightWhite   = ESC + "[97m"
)

func Fg256(color uint8) string {
	return ESC + "[38;5;" + strconv.FormatUint(uint64(color), 10) + "m"
}

func FgRgb(red, green, blue uint8) string {
	return ESC + "[38;2;" + strconv.FormatUint(uint64(red), 10) + ";" + strconv.FormatUint(uint64(green), 10) + ";" + strconv.FormatUint(uint64(blue), 10) + "m"
}

// Background Color
const (
	BgBlack   = ESC + "[40m"
	BgRed     = ESC + "[41m"
	BgGreen   = ESC + "[42m"
	BgYellow  = ESC + "[43m"
	BgBlue    = ESC + "[44m"
	BgMagenta = ESC + "[45m"
	BgCyan    = ESC + "[46m"
	BgWhite   = ESC + "[47m"

	BgBrightBlack   = ESC + "[100m"
	BgBrightRed     = ESC + "[101m"
	BgBrightGreen   = ESC + "[102m"
	BgBrightYellow  = ESC + "[103m"
	BgBrightBlue    = ESC + "[104m"
	BgBrightMagenta = ESC + "[105m"
	BgBrightCyan    = ESC + "[106m"
	BgBrightWhite   = ESC + "[107m"
)

func BgRgb(red, green, blue uint8) string {
	return ESC + "[48;2;" + strconv.FormatUint(uint64(red), 10) + ";" + strconv.FormatUint(uint64(green), 10) + ";" + strconv.FormatUint(uint64(blue), 10) + "m"
}

func Bg256(color uint8) string {
	return ESC + "[48;5;" + strconv.FormatUint(uint64(color), 10) + "m"
}

// Cursor Movement

func CursorMove(col, line int) {
	fmt.Print(ESC+"[", line, ";", col, "H")
}

func CursorUp(line int) {
	fmt.Print(ESC+"[", line, "A")
}

func CursorDown(line int) {
	fmt.Print(ESC+"[", line, "B")
}

func CursorRight(col int) {
	fmt.Print(ESC+"[", col, "C")
}

func CursorLeft(col int) {
	fmt.Print(ESC+"[", col, "D")
}

func CursorCol(col int) {
	fmt.Print(ESC+"[", col, "G")
}

func CursorHome() {
	fmt.Print(ESC + "[H")
}

func CursorSave() {
	fmt.Print(ESC + "[s")
}

func CursorRestore() {
	fmt.Print(ESC + "[u")
}

func CursorInvisible() {
	fmt.Print(ESC + "[?25l")
}

func CursorVisible() {
	fmt.Print(ESC + "[?25h")
}

func ScreenSave() {
	fmt.Print(ESC + "[?47h")
}

func ScreenRestore() {
	fmt.Print(ESC + "[?47l")
}

// Clear
const (
	ClearScreen     = ESC + "[2J"
	ClearScreenUp   = ESC + "[1J"
	ClearScreenDown = ESC + "[0J"
	ClearScreenEnd  = ESC + "[J"
	ClearLineStart  = ESC + "[1K"
	ClearLineEnd    = ESC + "[K"
	ClearLine       = ESC + "[2K"
)
