package tast

import (
	"go/ast"
	"go/token"
	"testing"
)

func NewPackage(t testing.TB, fset *token.FileSet, files map[string]*ast.File, importer ast.Importer, universe *ast.Scope) *ast.Package {
	t.Helper()
	r0, err := ast.NewPackage(fset, files, importer, universe)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
