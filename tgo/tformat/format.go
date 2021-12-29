package tformat

import (
	"go/format"
	"go/token"
	"io"
	"testing"
)

func Node(t testing.TB, dst io.Writer, fset *token.FileSet, node any) {
	t.Helper()
	err := format.Node(dst, fset, node)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func Source(t testing.TB, src []byte) []byte {
	t.Helper()
	r0, err := format.Source(src)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
