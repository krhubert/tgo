package tfs

import (
	"io/fs"
	"testing"
)

func Stat(t testing.TB, fsys fs.FS, name string) fs.FileInfo {
	t.Helper()
	r0, err := fs.Stat(fsys, name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
