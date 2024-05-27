package main

import (
	"fmt"

	"pat/ansi"
	"pat/img"
)

func main() {
	image, err := img.Decode("./img/test_img.png")
	if err != nil {
		panic(err)
	}

	image = img.Resize(image, 80, 0)

	previousPixels := ansi.PixelPair{}
	for y := 0; y < image.Bounds().Dy(); y += 2 {
		for x := 0; x < image.Bounds().Dx(); x++ {

			pixels := ansi.Pair(image.At(x, y+1), image.At(x, y))

			if pixels == previousPixels {
				print("▄")
				continue
			}

			fmt.Print(ansi.ConvertPixels(pixels))
			previousPixels = pixels
		}
		fmt.Println(ansi.Reset())
	}
}
