package tparser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"testing"
)

func ParseFile(t testing.TB, fset *token.FileSet, filename string, src any, mode parser.Mode) (f *ast.File) {
	t.Helper()
	r0, err := parser.ParseFile(fset, filename, src, mode)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseDir(t testing.TB, fset *token.FileSet, path string, filter func(fs.FileInfo) bool, mode parser.Mode) (pkgs map[string]*ast.Package) {
	t.Helper()
	r0, err := parser.ParseDir(fset, path, filter, mode)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseExprFrom(t testing.TB, fset *token.FileSet, filename string, src any, mode parser.Mode) (expr ast.Expr) {
	t.Helper()
	r0, err := parser.ParseExprFrom(fset, filename, src, mode)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseExpr(t testing.TB, x string) ast.Expr {
	t.Helper()
	r0, err := parser.ParseExpr(x)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
