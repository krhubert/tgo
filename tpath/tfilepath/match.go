package tfilepath

import (
	"path/filepath"
	"testing"
)

func Match(t testing.TB, pattern, name string) (matched bool) {
	t.Helper()
	r0, err := filepath.Match(pattern, name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Glob(t testing.TB, pattern string) (matches []string) {
	t.Helper()
	r0, err := filepath.Glob(pattern)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
