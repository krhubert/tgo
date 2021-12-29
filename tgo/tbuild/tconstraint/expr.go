package tconstraint

import (
	"go/build/constraint"
	"testing"
)

func Parse(t testing.TB, line string) constraint.Expr {
	t.Helper()
	r0, err := constraint.Parse(line)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func PlusBuildLines(t testing.TB, x constraint.Expr) []string {
	t.Helper()
	r0, err := constraint.PlusBuildLines(x)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
