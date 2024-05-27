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

	ansi.PrintImage(image, 10, 0)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
