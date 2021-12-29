package tfs

import (
	"io/fs"
	"testing"
)

func ReadFile(t testing.TB, fsys fs.FS, name string) []byte {
	t.Helper()
	r0, err := fs.ReadFile(fsys, name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
