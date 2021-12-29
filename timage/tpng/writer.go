package tpng

import (
	"image"
	"image/png"
	"io"
	"testing"
)

func Encode(t testing.TB, w io.Writer, m image.Image) {
	t.Helper()
	err := png.Encode(w, m)
	if err != nil {
		t.Fatal(err)
	}
	return
}
