package ttypes

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func Eval(t testing.TB, fset *token.FileSet, pkg *types.Package, pos token.Pos, expr string) (_ types.TypeAndValue) {
	t.Helper()
	r0, err := types.Eval(fset, pkg, pos, expr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func CheckExpr(t testing.TB, fset *token.FileSet, pkg *types.Package, pos token.Pos, expr ast.Expr, info *types.Info) {
	t.Helper()
	err := types.CheckExpr(fset, pkg, pos, expr, info)
	if err != nil {
		t.Fatal(err)
	}
	return
}
