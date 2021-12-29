package timage

import (
	"image"
	"io"
	"testing"
)

func Decode(t testing.TB, r io.Reader) (image.Image, string) {
	t.Helper()
	r0, r1, err := image.Decode(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1
}

func DecodeConfig(t testing.TB, r io.Reader) (image.Config, string) {
	t.Helper()
	r0, r1, err := image.DecodeConfig(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1
}
