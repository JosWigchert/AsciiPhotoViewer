package main

import (
	"ascii/jutils/image"
	"ascii/jutils/terminal"
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("program.prof")
	if err != nil {
		fmt.Println("Error creating profile file:", err)
		return
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	img, err := image.GetImageFromFilePath("C:\\Users\\starw\\Pictures\\Screenshots\\Screenshot 2024-04-04 223504.png")
	// img, err := image.GetImageFromFilePath("C:\\Users\\starw\\Pictures\\Screenshots\\Screenshot 2024-03-31 140256.png")
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
	convertOptions := image.ConvertOptions{
		Color: true,
	}

	terminal.MoveCursorToStart()
	fmt.Print(image.ConvertToASCII(img, convertOptions))
}
