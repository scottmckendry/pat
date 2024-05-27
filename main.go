package main

import (
	"fmt"
	"os"

	"pat/ansi"
	"pat/img"
)

func main() {
	path := os.Args[1]
	if !pathExists(path) {
		fmt.Println("Could not find file:", path)
		os.Exit(1)
	}

	image, err := img.Decode(path)
	if err != nil {
		fmt.Println("Could not decode image:", err)
		os.Exit(1)
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

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
