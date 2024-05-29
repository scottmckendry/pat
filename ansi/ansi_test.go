package ansi_test

import (
	"bytes"
	"image/color"
	"os"
	"testing"

	"github.com/scottmckendry/pat/ansi"
	"github.com/scottmckendry/pat/img"
)

func TestPair(t *testing.T) {
	color1 := color.Color(color.RGBA{0, 0, 0, 0})
	color2 := color.Color(color.RGBA{255, 255, 255, 255})

	pixels := ansi.Pair(color1, color2)

	want := ansi.PixelPair{color1, color2}
	if pixels != want {
		t.Fatalf("Expected %v, got %v", want, pixels)
	}
}

func TestConvertPixels(t *testing.T) {
	pixels := ansi.PixelPair{
		color.RGBA{0, 100, 200, 255},
		color.RGBA{5, 25, 55, 155},
	}

	want := "\033[38;2;0;100;200;48;2;5;25;55m▄"
	got := ansi.ConvertPixels(pixels)

	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}

func TestReset(t *testing.T) {
	want := "\033[0m"
	got := ansi.Reset()

	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}

func TestPrintPixels(t *testing.T) {
	pixels := ansi.PixelPair{
		color.RGBA{0, 100, 200, 255},
		color.RGBA{5, 25, 55, 155},
	}

	want := "\033[38;2;0;100;200;48;2;5;25;55m▄"
	got := ansi.ConvertPixels(pixels)

	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}

func TestPrintImage(t *testing.T) {
	// read from testdata file. This is the expected output when the test image is rendered at a max width of 10.
	want, err := os.ReadFile("./testdata/img_out_10.txt")
	if err != nil {
		t.Fatalf("Could not read from testdata file: %v", err)
	}

	// remove windows line endings if present
	want = bytes.ReplaceAll(want, []byte("\r\n"), []byte("\n"))

	// keep a backup of the real stdout, and replace it with a pipe.
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// render the image
	img, err := img.Decode("../img/test_img.png")
	if err != nil {
		t.Fatalf("Could not decode image: %v", err)
	}

	ansi.PrintImage(img, 10, 0)

	// restore the real stdout
	_ = w.Close()
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	os.Stdout = old

	if buf.String() != string(want) {
		t.Fatalf("Expected %q, got %q", want, buf.String())
	}
}
