package tmacho

import (
	"debug/macho"
	"io"
	"testing"
)

func Open(t testing.TB, name string) *macho.File {
	t.Helper()
	r0, err := macho.Open(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func NewFile(t testing.TB, r io.ReaderAt) *macho.File {
	t.Helper()
	r0, err := macho.NewFile(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
