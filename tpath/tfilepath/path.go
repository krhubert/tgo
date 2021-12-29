package tfilepath

import (
	"io/fs"
	"path/filepath"
	"testing"
)

func EvalSymlinks(t testing.TB, path string) string {
	t.Helper()
	r0, err := filepath.EvalSymlinks(path)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Abs(t testing.TB, path string) string {
	t.Helper()
	r0, err := filepath.Abs(path)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Rel(t testing.TB, basepath, targpath string) string {
	t.Helper()
	r0, err := filepath.Rel(basepath, targpath)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func WalkDir(t testing.TB, root string, fn fs.WalkDirFunc) {
	t.Helper()
	err := filepath.WalkDir(root, fn)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func Walk(t testing.TB, root string, fn filepath.WalkFunc) {
	t.Helper()
	err := filepath.Walk(root, fn)
	if err != nil {
		t.Fatal(err)
	}
	return
}
