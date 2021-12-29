package tgzip

import (
	"compress/gzip"
	"io"
	"testing"
)

func NewWriterLevel(t testing.TB, w io.Writer, level int) *gzip.Writer {
	t.Helper()
	r0, err := gzip.NewWriterLevel(w, level)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
