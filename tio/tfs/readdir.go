package tfs

import (
	"io/fs"
	"testing"
)

func ReadDir(t testing.TB, fsys fs.FS, name string) []fs.DirEntry {
	t.Helper()
	r0, err := fs.ReadDir(fsys, name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
