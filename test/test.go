package main

import (
	"fmt"
	"github.com/cloudyhadi/ansi"
)

func HelloJohn() {
	var name string = "John"
	ansi.ClearScreen()
	ansi.CursorHome()
	fmt.Println(ansi.White + "Your name is " + ansi.Blue + name)
	fmt.Println(ansi.White + "Your name is " + ansi.BrightRed + name)
	fmt.Println(ansi.White + "Your name is " + ansi.BgWhite + ansi.Black + name)

	fmt.Println(ansi.Reset + "Your name is " + ansi.FgRgb(100, 200, 200) + name)
}

func main() {
	var name string

	ansi.CursorSave()
	ansi.ScreenSave()
	ansi.CursorHome()
	fmt.Print(ansi.BgBlue + ansi.Black + "Hello world!" + ansi.Reset)
	ansi.CursorMove(0, 1)
	fmt.Print("What's your name ?" + ansi.BrightCyan)
	ansi.CursorDown(1)
	ansi.CursorCol(0)
	fmt.Scan(&name)
	ansi.CursorRestore()
	ansi.ScreenRestore()

	fmt.Println(ansi.Reset + "Your name is " + ansi.FgRgb(200, 100, 100) + name)
	fmt.Println(ansi.Reset + "Your name is " + ansi.FgRgb(100, 200, 100) + name)
	fmt.Println(ansi.Reset + "Your name is " + ansi.FgRgb(100, 200, 200) + name)
	fmt.Println(ansi.Reset + "Your name is " + ansi.FgRgb(150, 150, 150) + name)
	fmt.Println(ansi.Reset + "Your name is " + ansi.FgRgb(255, 0, 0) + name)
	fmt.Println(ansi.Reset + "Your name is " + ansi.FgRgb(0, 255, 0) + name)
	fmt.Println(ansi.Reset + "Your name is " + ansi.FgRgb(0, 0, 255) + name)
	HelloJohn()
}
