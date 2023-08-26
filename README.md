<p align="center">
  <img width="60%" src="https://cdn.nostr.build/i/2c5ce0ee4f2a63742d357ff060a21dd1896e9fcdc85a1784ed7d22706330a94e.png" />
</p>

# ANSI
### Go Package for Colorful Terminal Output and Control

The 'ansi' package provides a simple and user-friendly way to add color and formatting to terminal output in Go applications. 
Enhance your command-line interfaces with vibrant text, background colors, and text styles, making your output more readable and visually appealing.

Features:

- Easy-to-use API for applying ANSI escape codes
- Support for a variety of text colors and background colors
- Different text styles like bold, underline, and more
- Control the cursor position for dynamic output layouts
- Clear the terminal screen for cleaner interfaces
- Compatible with various terminal emulators and platforms

---

## Usage

### Installation:
```bash
go get -u github.com/anotherhadi/ansi@latest
```

### Importing:
```go
import (
  "fmt"
  "github.com/anotherhadi/ansi"
)
```

### Basic Usage:
```go
func HelloJohn() {
  var name string = "John"
  ansi.ClearScreen()
  ansi.CursorHome()
  fmt.Println(ansi.White + "Your name is " + ansi.Blue + name)
  fmt.Println(ansi.White + "Your name is " + ansi.BrightRed + name)
  fmt.Println(ansi.White + "Your name is " + ansi.BgWhite + ansi.Black + name)

  fmt.Println(ansi.Reset + "Your name is " + ansi.FgRgb(100, 200, 200) + name)
}
```
Example available in "test/test.go"

---
Reference: [ANSI.md](https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797#file-ansi-md)

<img src="https://img.buymeacoffee.com/button-api/?text=Buy me a cookie&emoji=🍪&slug=anotherhadi&button_colour=eed2cc&font_colour=000000&font_family=Inter&outline_colour=ffffff&coffee_colour=ff0000" />
