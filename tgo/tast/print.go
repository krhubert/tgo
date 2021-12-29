package tast

import (
	"go/ast"
	"go/token"
	"io"
	"testing"
)

func Fprint(t testing.TB, w io.Writer, fset *token.FileSet, x any, f ast.FieldFilter) {
	t.Helper()
	err := ast.Fprint(w, fset, x, f)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func Print(t testing.TB, fset *token.FileSet, x any) {
	t.Helper()
	err := ast.Print(fset, x)
	if err != nil {
		t.Fatal(err)
	}
	return
}
