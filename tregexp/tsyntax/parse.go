package tsyntax

import (
	"regexp/syntax"
	"testing"
)

func Parse(t testing.TB, s string, flags syntax.Flags) *syntax.Regexp {
	t.Helper()
	r0, err := syntax.Parse(s, flags)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
