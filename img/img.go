package img

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/nfnt/resize"
)

// Decode reads an image from a file and returns the image.Image.
// If the file cannot be read, an error is returned.
func Decode(filepath string) (image.Image, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// Resize resizes an image to the specified width and height.
// If either width or height is 0, the aspect ratio is maintained.
func Resize(image image.Image, width, height int) image.Image {
	return resize.Resize(uint(width), uint(height), image, resize.Lanczos3)
}
