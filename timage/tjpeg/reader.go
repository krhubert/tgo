package tjpeg

import (
	"image"
	"image/jpeg"
	"io"
	"testing"
)

func Decode(t testing.TB, r io.Reader) image.Image {
	t.Helper()
	r0, err := jpeg.Decode(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DecodeConfig(t testing.TB, r io.Reader) image.Config {
	t.Helper()
	r0, err := jpeg.DecodeConfig(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
