package img

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"

	"github.com/nfnt/resize"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

// Decode reads an image from a file and returns the image.Image.
// If the file cannot be read, an error is returned.
func Decode(filepath string) (image.Image, error) {
	if isUrl(filepath) {
		return decodeFromUrl(filepath)
	}

	if !pathExists(filepath) {
		return nil, os.ErrNotExist
	}

	return decodeFromFile(filepath)
}

// Resizes an image to the specified width and height.
// If either width or height is 0, the aspect ratio is maintained.
// If the output of the resize has an uneven height, a final resize is done to make it even.
// This prevents a line of black pixels from appearing at the bottom of the image when rendered.
func Resize(image image.Image, width, height int) image.Image {
	image = resize.Resize(uint(width), uint(height), image, resize.Lanczos3)

	// check if the height is even and resize again if it is not
	if image.Bounds().Dy()%2 != 0 {
		image = resize.Resize(
			uint(image.Bounds().Dx()),
			uint(image.Bounds().Dy()+1),
			image,
			resize.Lanczos3,
		)
	}

	return image
}

// execute a GET request to the provided URL and decode the image
func decodeFromUrl(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// open a file and decode the image
func decodeFromFile(filepath string) (image.Image, error) {
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

// check if the path is a URL
func isUrl(path string) bool {
	return path[:4] == "http"
}

// check if a file exists at the provided path
func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
