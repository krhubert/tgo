package tfs

import (
	"io/fs"
	"testing"
)

func Sub(t testing.TB, fsys fs.FS, dir string) fs.FS {
	t.Helper()
	r0, err := fs.Sub(fsys, dir)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
