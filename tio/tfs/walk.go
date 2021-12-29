package tfs

import (
	"io/fs"
	"testing"
)

func WalkDir(t testing.TB, fsys fs.FS, root string, fn fs.WalkDirFunc) {
	t.Helper()
	err := fs.WalkDir(fsys, root, fn)
	if err != nil {
		t.Fatal(err)
	}
	return
}
