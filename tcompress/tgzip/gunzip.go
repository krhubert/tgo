package tgzip

import (
	"compress/gzip"
	"io"
	"testing"
)

func NewReader(t testing.TB, r io.Reader) *gzip.Reader {
	t.Helper()
	r0, err := gzip.NewReader(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
