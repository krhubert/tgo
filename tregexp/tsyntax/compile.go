package tsyntax

import (
	"regexp/syntax"
	"testing"
)

func Compile(t testing.TB, re *syntax.Regexp) *syntax.Prog {
	t.Helper()
	r0, err := syntax.Compile(re)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
