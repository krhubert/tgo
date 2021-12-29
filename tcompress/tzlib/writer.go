package tzlib

import (
	"compress/zlib"
	"io"
	"testing"
)

func NewWriterLevel(t testing.TB, w io.Writer, level int) *zlib.Writer {
	t.Helper()
	r0, err := zlib.NewWriterLevel(w, level)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func NewWriterLevelDict(t testing.TB, w io.Writer, level int, dict []byte) *zlib.Writer {
	t.Helper()
	r0, err := zlib.NewWriterLevelDict(w, level, dict)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
