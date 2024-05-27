package img_test

import (
	"testing"

	"pat/img"
)

func TestDecode(t *testing.T) {
	_, err := img.Decode("./test_img.png")
	if err != nil {
		t.Fatalf("Error decoding image: %v", err)
	}
}

func TestDecodeBadPath(t *testing.T) {
	_, err := img.Decode("./bad_path.png")
	if err == nil {
		t.Fatalf("Expected error decoding image")
	}
}

func TestResize(t *testing.T) {
	image, err := img.Decode("./test_img.png")
	if err != nil {
		t.Fatalf("Error decoding image: %v", err)
	}

	image = img.Resize(image, 80, 0)
	if image.Bounds().Dx() != 80 {
		t.Fatalf("Expected width to be 80, got %d", image.Bounds().Dx())
	}
}
