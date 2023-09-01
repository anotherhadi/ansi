// https://github.com/anotherhadi/ansi
package ansi

import (
	"fmt"
	"strconv"
)

const ESC = "\033"

// Text Styling
// Ex: fmt.Println(ansi.Underline+ "Underlined text")
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
// Ex: fmt.Println(ansi.Blue+ "Blue text")
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

// 256 Text Color
// Ex: fmt.Println(ansi.Fg256(27)+ "Blue text")
func Fg256(color uint8) string {
	return ESC + "[38;5;" + strconv.FormatUint(uint64(color), 10) + "m"
}

// RGB Text Color
// Ex: fmt.Println(ansi.FgRgb(30, 30, 255)+ "Blue text")
func FgRgb(red, green, blue uint8) string {
	return ESC + "[38;2;" + strconv.FormatUint(uint64(red), 10) + ";" + strconv.FormatUint(uint64(green), 10) + ";" + strconv.FormatUint(uint64(blue), 10) + "m"
}

// Background Color
// Ex: fmt.Println(ansi.BgBlue+ "Text with a blue background")
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

// 256 Background Color
// Ex: fmt.Println(ansi.Bg256(27)+ "Text with a blue background")
func BgRgb(red, green, blue uint8) string {
	return ESC + "[48;2;" + strconv.FormatUint(uint64(red), 10) + ";" + strconv.FormatUint(uint64(green), 10) + ";" + strconv.FormatUint(uint64(blue), 10) + "m"
}

// RGB Background Color
// Ex: fmt.Println(ansi.BgRgb(30, 30, 255)+ "Text with a blue background")
func Bg256(color uint8) string {
	return ESC + "[48;5;" + strconv.FormatUint(uint64(color), 10) + "m"
}

//// Cursor Movement

// Move cursor to {line}, {col}
// Ex: ansi.CursorMove(10,20)
func CursorMove(line, col int) {
	fmt.Print(ESC+"[", line, ";", col, "H")
}

// Move cursor {line} up
// Ex: ansi.CursorUp(10)
func CursorUp(line int) {
	fmt.Print(ESC+"[", line, "A")
}

// Move cursor {line} down
// Ex: ansi.CursorDown(10)
func CursorDown(line int) {
	fmt.Print(ESC+"[", line, "B")
}

// Move cursor {col} right
// Ex: ansi.CursorRight(10)
func CursorRight(col int) {
	fmt.Print(ESC+"[", col, "C")
}

// Move cursor {col} left
// Ex: ansi.CursorLeft(10)
func CursorLeft(col int) {
	fmt.Print(ESC+"[", col, "D")
}

// Move cursor to {col}
// Ex: ansi.CursorCol(10)
func CursorCol(col int) {
	fmt.Print(ESC+"[", col, "G")
}

// Move cursor to Home (0,0)
// Ex: ansi.CursorHome()
func CursorHome() {
	fmt.Print(ESC + "[H")
}

// Save cursor position
func CursorSave() {
	fmt.Print(ESC + "[s")
}

// Restore cursor position
func CursorRestore() {
	fmt.Print(ESC + "[u")
}

// Make cursor invisible
func CursorInvisible() {
	fmt.Print(ESC + "[?25l")
}

// Make cursor visible
func CursorVisible() {
	fmt.Print(ESC + "[?25h")
}

// Save the screen
func ScreenSave() {
	fmt.Print(ESC + "[?47h")
}

// Restore the screen
func ScreenRestore() {
	fmt.Print(ESC + "[?47l")
}

//// Clear/Erase

// Clear the full screen
func ClearScreen() {
	fmt.Print(ESC + "[2J")
}

// Clear screen up
func ClearScreenUp() {
	fmt.Print(ESC + "[1J")
}

// Clear screen down
func ClearScreenDown() {
	fmt.Print(ESC + "[0J")
}

// Clear screen to the end
func ClearScreenEnd() {
	fmt.Print(ESC + "[J")
}

// Clear the entire line
func ClearLine() {
	fmt.Print(ESC + "[2K")
}

// Clear start of line to the cursor
func ClearLineStart() {
	fmt.Print(ESC + "[1K")
}

// Clear from cursor to end of line
func ClearLineEnd() {
	fmt.Print(ESC + "[K")
}
