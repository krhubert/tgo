package tprinter

import (
	"go/printer"
	"go/token"
	"io"
	"testing"
)

func Fprint(t testing.TB, output io.Writer, fset *token.FileSet, node any) {
	t.Helper()
	err := printer.Fprint(output, fset, node)
	if err != nil {
		t.Fatal(err)
	}
	return
}
