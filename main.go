package main

import (
	"flag"
	"fmt"
	"os"

	"pat/ansi"
	"pat/img"
)

func main() {
	flag.Usage = func() {
		fmt.Print("Usage: pat <path-to-image> [options]\n\n")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	width := flag.Int("w", 100, "Number of columns to use for the image\n  ")
	height := flag.Int(
		"h",
		0,
		"Number of rows to use for the image. Using 0 in either width or height will\n"+"preserve the aspect ratio of the image.\n   (default 0)",
	)
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

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

	ansi.PrintImage(image, *width, *height)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
