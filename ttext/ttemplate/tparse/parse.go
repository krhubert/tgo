package tparse

import (
	"testing"
	"text/template/parse"
)

func Parse(t testing.TB, name, text, leftDelim, rightDelim string, funcs ...map[string]any) map[string]*parse.Tree {
	t.Helper()
	r0, err := parse.Parse(name, text, leftDelim, rightDelim, funcs...)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
