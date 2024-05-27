package ansi_test

import (
	"image/color"
	"testing"

	"pat/ansi"
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
