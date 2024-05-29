package ansi

import (
	"fmt"
	"image"
	"image/color"

	"github.com/scottmckendry/pat/img"
)

// PixelPair represents a pair of pixels, one for the top half of a character and one for the bottom half.
// This accounts for terminal characters being twice as tall as they are wide.
type PixelPair struct {
	TopPixel    color.Color
	BottomPixel color.Color
}

// Converts a PixelPair to an ANSI escape code that will render the pair of pixels as a single character.
func ConvertPixels(pixels PixelPair) string {
	fgR, fgG, fgB, fgA := pixels.TopPixel.RGBA()
	bgR, bgG, bgB, bgA := pixels.BottomPixel.RGBA()

	transparentThreshold := uint32(200)

	if fgA < transparentThreshold && bgA < transparentThreshold {
		return "\033[0m "
	}

	if bgA < transparentThreshold {
		return fmt.Sprintf("\033[38;2;%d;%d;%dm▄", fgR/256, fgG/256, fgB/256)
	}

	if fgA < transparentThreshold {
		// Use a [[reverse video]] character if the top pixel is transparent.
		return fmt.Sprintf("\033[7;38;2;%d;%d;%dm▄\033[0m", bgR/256, bgG/256, bgB/256)
	}

	return fmt.Sprintf(
		"\033[38;2;%d;%d;%d;48;2;%d;%d;%dm▄",
		fgR/256,
		fgG/256,
		fgB/256,
		bgR/256,
		bgG/256,
		bgB/256,
	)
}

// Pair creates a PixelPair from two color.Colors.
func Pair(topPixel color.Color, bottomPixel color.Color) PixelPair {
	return PixelPair{topPixel, bottomPixel}
}

// Returns the ANSI escape code to reset the terminal colors.
func Reset() string {
	return "\033[0m"
}

func PrintPixels(pixels PixelPair) {
	fmt.Print(ConvertPixels(pixels))
}

func PrintImage(image image.Image, width int, height int) {
	image = img.Resize(image, width, height)
	for y := 0; y < image.Bounds().Dy(); y += 2 {
		for x := 0; x < image.Bounds().Dx(); x++ {
			pixels := Pair(image.At(x, y+1), image.At(x, y))
			PrintPixels(pixels)
		}
		fmt.Println(Reset())
	}
}
