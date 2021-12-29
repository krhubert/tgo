package tregexp

import (
	"io"
	"regexp"
	"testing"
)

func Compile(t testing.TB, expr string) *regexp.Regexp {
	t.Helper()
	r0, err := regexp.Compile(expr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func CompilePOSIX(t testing.TB, expr string) *regexp.Regexp {
	t.Helper()
	r0, err := regexp.CompilePOSIX(expr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func MatchReader(t testing.TB, pattern string, r io.RuneReader) (matched bool) {
	t.Helper()
	r0, err := regexp.MatchReader(pattern, r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func MatchString(t testing.TB, pattern string, s string) (matched bool) {
	t.Helper()
	r0, err := regexp.MatchString(pattern, s)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Match(t testing.TB, pattern string, b []byte) (matched bool) {
	t.Helper()
	r0, err := regexp.Match(pattern, b)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
