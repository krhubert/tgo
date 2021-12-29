package tpng

import (
	"image"
	"image/png"
	"io"
	"testing"
)

func Decode(t testing.TB, r io.Reader) image.Image {
	t.Helper()
	r0, err := png.Decode(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DecodeConfig(t testing.TB, r io.Reader) image.Config {
	t.Helper()
	r0, err := png.DecodeConfig(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
