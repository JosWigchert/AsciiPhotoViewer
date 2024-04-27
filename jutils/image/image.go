package image

import (
	"image"
	"os"

	"github.com/nfnt/resize"
)

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

func ConvertToASCII(img image.Image) string {
	// gradient := []rune(" .:-=+*#%@")
	gradient := []rune(" `.-':_,^=;><+!rc*/z?sLTv)J7(|Fi{C}fI31tlu[neoZ5Yxjya]2ESwqkP6h9d4VpOGbUAKXHm8RD#$Bg0MNWQ%&@")
	gradientLength := uint32(len(gradient)) - 1

	bounds := img.Bounds()
	dy := bounds.Dy()
	dx := bounds.Dx()

	asciiImage := make([]rune, dx*dy+dy)

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gradientIndex := (r + g + b) / 3 / (0xffff / gradientLength)
			character := gradient[gradientIndex]
			asciiImage[(y*dx)+x+y] = character
		}
		asciiImage[(y*dx)+dx+y] = '\n'
	}

	return string(asciiImage)
}
