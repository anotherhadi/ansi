package ansi

import (
	"fmt"
)

func test() {
	var name string

	CursorSave()
	ScreenSave()
	CursorMove(10, 10)
	fmt.Print(BgBlue, Black, "Hello world!", Reset)
	CursorMove(10, 20)
	fmt.Print("What's your name ?", BrightCyan)
	CursorDown(1)
	CursorCol(10)
	fmt.Scan(&name)
	ScreenRestore()
	CursorRestore()
	fmt.Print(Reset, "Your name is ", FgRgb(0, 0, 255), name)

}
