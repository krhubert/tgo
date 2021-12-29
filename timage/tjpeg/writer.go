package tjpeg

import (
	"image"
	"image/jpeg"
	"io"
	"testing"
)

func Encode(t testing.TB, w io.Writer, m image.Image, o *jpeg.Options) {
	t.Helper()
	err := jpeg.Encode(w, m, o)
	if err != nil {
		t.Fatal(err)
	}
	return
}
