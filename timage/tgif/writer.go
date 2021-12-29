package tgif

import (
	"image"
	"image/gif"
	"io"
	"testing"
)

func EncodeAll(t testing.TB, w io.Writer, g *gif.GIF) {
	t.Helper()
	err := gif.EncodeAll(w, g)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func Encode(t testing.TB, w io.Writer, m image.Image, o *gif.Options) {
	t.Helper()
	err := gif.Encode(w, m, o)
	if err != nil {
		t.Fatal(err)
	}
	return
}
