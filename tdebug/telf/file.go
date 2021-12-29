package telf

import (
	"debug/elf"
	"io"
	"testing"
)

func Open(t testing.TB, name string) *elf.File {
	t.Helper()
	r0, err := elf.Open(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func NewFile(t testing.TB, r io.ReaderAt) *elf.File {
	t.Helper()
	r0, err := elf.NewFile(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
