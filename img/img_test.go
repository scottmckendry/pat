package img

import (
	"testing"
)

func TestDecode(t *testing.T) {
	_, err := Decode("./test_img.png")
	if err != nil {
		t.Fatalf("Error decoding image: %v", err)
	}
}

func TestDecodeBadPath(t *testing.T) {
	_, err := Decode("./bad_path.png")
	if err == nil {
		t.Fatalf("Expected error decoding image")
	}
}

func TestResize(t *testing.T) {
	image, err := Decode("./test_img.png")
	if err != nil {
		t.Fatalf("Error decoding image: %v", err)
	}

	image = Resize(image, 80, 0)
	if image.Bounds().Dx() != 80 {
		t.Fatalf("Expected width to be 80, got %d", image.Bounds().Dx())
	}
}

func TestDecodeFromUrl(t *testing.T) {
	_, err := decodeFromUrl("https://scottmckendry.tech/img/logo.webp")
	if err != nil {
		t.Fatalf("Error decoding image from URL: %v", err)
	}
}

func TestDecodeFromUrlBadUrl(t *testing.T) {
	_, err := decodeFromUrl("https://scottmckendry.tech/img/bad_url.webp")
	if err == nil {
		t.Fatalf("Expected error decoding image from URL")
	}
}

func TestDecodeFromFile(t *testing.T) {
	_, err := decodeFromFile("./test_img.png")
	if err != nil {
		t.Fatalf("Error decoding image from file: %v", err)
	}
}

func TestDecodeFromFileBadPath(t *testing.T) {
	_, err := decodeFromFile("./bad_path.png")
	if err == nil {
		t.Fatalf("Expected error decoding image from file")
	}
}

func TestIsUrl(t *testing.T) {
	if !isUrl("http://example.com") {
		t.Fatalf("Expected URL to be recognized")
	}
}

func TestIsUrlBadUrl(t *testing.T) {
	if isUrl("example.com") {
		t.Fatalf("Expected URL to not be recognized")
	}
}

func TestPathExists(t *testing.T) {
	if !pathExists("./test_img.png") {
		t.Fatalf("Expected path to exist")
	}
}

func TestPathExistsBadPath(t *testing.T) {
	if pathExists("./bad_path.png") {
		t.Fatalf("Expected path to not exist")
	}
}
