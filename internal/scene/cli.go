package scene

import "fmt"

func ClearCanvas() {
	fmt.Print("\033[2J\033[H") // clear
}

func HideCursor() {
	fmt.Print("\033[?25l")
}

func ShowCursor() {
	fmt.Print("\033[?25h")
}

func MoveCursor(x uint8, y uint8) {
	fmt.Printf("\033[%d;%dH", y+1, x+1)
}

func WriteAt(x uint8, y uint8, str string) {
	MoveCursor(x, y)
	fmt.Print(str)

}

func DrawCanvas(str string) {
	fmt.Print(str)
}
