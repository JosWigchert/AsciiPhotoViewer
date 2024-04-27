package image

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/nfnt/resize"
)

type ConvertOptions struct {
	Color bool
}

// DefaultConvertOptions is the default options for converting to ASCII art
var DefaultConvertOptions = ConvertOptions{
	Color: false,
}

func GetImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

func ResizeImage(img image.Image, width, height int) image.Image {
	resizedImg := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
	return resizedImg
}

func ConvertToASCII(img image.Image, options ConvertOptions) string {
	gradient := []rune(" .:-+*#%@")
	// gradient := []rune(" `.-':_,^=;><+!rc*/z?sLTv)J7(|Fi{C}fI31tlu[neoZ5Yxjya]2ESwqkP6h9d4VpOGbUAKXHm8RD#$Bg0MNWQ%&@")
	if options.Color {
		gradient = []rune("@@")
	}

	gradientLength := uint32(len(gradient)) - 1

	bounds := img.Bounds()
	dy := bounds.Dy()
	dx := bounds.Dx()

	asciiImageLength := dx*dy + dy
	if options.Color {
		asciiImageLength += asciiImageLength * 19
	}
	asciiImage := make([]rune, asciiImageLength)

	currentIndex := 0
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gradientIndex := (r + g + b) / 3 / (0xffff / gradientLength)
			character := gradient[gradientIndex]

			if options.Color {
				colorLength := uint32(0xffff / 254)
				rgb := []rune(fmt.Sprintf("\033[38;2;%d;%d;%dm", r/colorLength, g/colorLength, b/colorLength))
				for _, v := range rgb {
					asciiImage[currentIndex] = v
					currentIndex += 1
				}
			}

			asciiImage[currentIndex] = character
			currentIndex += 1
		}
		asciiImage[currentIndex] = '\n'
		currentIndex += 1
	}

	if options.Color {
		rgb := []rune("\033[0m")
		for _, v := range rgb {
			asciiImage[currentIndex] = v
			currentIndex += 1
		}
	}

	return string(asciiImage)
}
