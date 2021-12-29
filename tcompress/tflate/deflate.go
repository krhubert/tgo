package tflate

import (
	"compress/flate"
	"io"
	"testing"
)

func NewWriter(t testing.TB, w io.Writer, level int) *flate.Writer {
	t.Helper()
	r0, err := flate.NewWriter(w, level)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func NewWriterDict(t testing.TB, w io.Writer, level int, dict []byte) *flate.Writer {
	t.Helper()
	r0, err := flate.NewWriterDict(w, level, dict)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
