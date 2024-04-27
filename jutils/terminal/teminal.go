package terminal

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func GetTerminalSize() (width, height int, err error) {
	var fd = int(os.Stdout.Fd())
	width, height, err = term.GetSize(fd)
	if err != nil {
		return 0, 0, err
	}
	return width, height, nil
}

func MoveCursorToStart() {
	fmt.Print("\033[H")  // ANSI escape code for moving cursor to the top-left corner
	fmt.Print("\033[#B") // ANSI escape code for moving cursor to the top-left corner
}
