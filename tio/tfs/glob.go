package tfs

import (
	"io/fs"
	"testing"
)

func Glob(t testing.TB, fsys fs.FS, pattern string) (matches []string) {
	t.Helper()
	r0, err := fs.Glob(fsys, pattern)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
