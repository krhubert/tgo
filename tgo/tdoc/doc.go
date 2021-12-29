package tdoc

import (
	"go/ast"
	"go/doc"
	"go/token"
	"testing"
)

func NewFromFiles(t testing.TB, fset *token.FileSet, files []*ast.File, importPath string, opts ...any) *doc.Package {
	t.Helper()
	r0, err := doc.NewFromFiles(fset, files, importPath, opts)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
