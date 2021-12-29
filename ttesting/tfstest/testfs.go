package tfstest

import (
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestFS(t testing.TB, fsys fs.FS, expected ...string) {
	t.Helper()
	err := fstest.TestFS(fsys, expected...)
	if err != nil {
		t.Fatal(err)
	}
	return
}
