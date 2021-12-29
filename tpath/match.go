package tpath

import (
	"path"
	"testing"
)

func Match(t testing.TB, pattern, name string) (matched bool) {
	t.Helper()
	r0, err := path.Match(pattern, name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
