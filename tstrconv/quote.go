package tstrconv

import (
	"strconv"
	"testing"
)

func UnquoteChar(t testing.TB, s string, quote byte) (value rune, multibyte bool, tail string) {
	t.Helper()
	r0, r1, r2, err := strconv.UnquoteChar(s, quote)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1, r2
}

func QuotedPrefix(t testing.TB, s string) string {
	t.Helper()
	r0, err := strconv.QuotedPrefix(s)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Unquote(t testing.TB, s string) string {
	t.Helper()
	r0, err := strconv.Unquote(s)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
