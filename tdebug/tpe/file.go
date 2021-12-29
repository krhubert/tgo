package tpe

import (
	"debug/pe"
	"io"
	"testing"
)

func Open(t testing.TB, name string) *pe.File {
	t.Helper()
	r0, err := pe.Open(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func NewFile(t testing.TB, r io.ReaderAt) *pe.File {
	t.Helper()
	r0, err := pe.NewFile(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
