package tplan9obj

import (
	"debug/plan9obj"
	"io"
	"testing"
)

func Open(t testing.TB, name string) *plan9obj.File {
	t.Helper()
	r0, err := plan9obj.Open(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func NewFile(t testing.TB, r io.ReaderAt) *plan9obj.File {
	t.Helper()
	r0, err := plan9obj.NewFile(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
