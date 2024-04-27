package main

import (
	"ascii/jutils/image"
	"ascii/jutils/terminal"
	"fmt"
	"os"
	"runtime/pprof"

	_ "image/jpeg"
	_ "image/png"
)

func main() {
	f, err := os.Create("program.prof")
	if err != nil {
		fmt.Println("Error creating profile file:", err)
		return
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	img, err := image.GetImageFromFilePath("C:\\Users\\starw\\Pictures\\Screenshots\\Screenshot 2023-12-09 191454.png")
	if err != nil {
		fmt.Println("Error reading image:", err)
		return
	}

	width, height, err := terminal.GetTerminalSize()
	if err != nil {
		fmt.Println("Error reading terminal size:", err)
		return
	}

	img = image.ResizeImage(img, width, height-1)

	// for i := 0; i < 100; i++ {
	// 	terminal.MoveCursorToStart()
	// 	fmt.Print(image.ConvertToASCII(img))
	// }

	fmt.Printf("runes in \033[38;2;1;255;20mrgb\033[0m %d\n", len([]rune("\033[38;2;20;90;200m")))
}
