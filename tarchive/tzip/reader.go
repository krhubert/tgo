package tzip

import (
	"archive/zip"
	"io"
	"testing"
)

func OpenReader(t testing.TB, name string) *zip.ReadCloser {
	t.Helper()
	r0, err := zip.OpenReader(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func NewReader(t testing.TB, r io.ReaderAt, size int64) *zip.Reader {
	t.Helper()
	r0, err := zip.NewReader(r, size)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
