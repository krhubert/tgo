package tmacho

import (
	"debug/macho"
	"io"
	"testing"
)

func NewFatFile(t testing.TB, r io.ReaderAt) *macho.FatFile {
	t.Helper()
	r0, err := macho.NewFatFile(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func OpenFat(t testing.TB, name string) *macho.FatFile {
	t.Helper()
	r0, err := macho.OpenFat(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
