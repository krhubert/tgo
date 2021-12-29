package tgif

import (
	"image"
	"image/gif"
	"io"
	"testing"
)

func Decode(t testing.TB, r io.Reader) image.Image {
	t.Helper()
	r0, err := gif.Decode(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DecodeAll(t testing.TB, r io.Reader) *gif.GIF {
	t.Helper()
	r0, err := gif.DecodeAll(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DecodeConfig(t testing.TB, r io.Reader) image.Config {
	t.Helper()
	r0, err := gif.DecodeConfig(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
