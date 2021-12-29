package tzlib

import (
	"compress/zlib"
	"io"
	"testing"
)

func NewReader(t testing.TB, r io.Reader) io.ReadCloser {
	t.Helper()
	r0, err := zlib.NewReader(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func NewReaderDict(t testing.TB, r io.Reader, dict []byte) io.ReadCloser {
	t.Helper()
	r0, err := zlib.NewReaderDict(r, dict)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
