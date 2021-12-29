package tioutil

import (
	"io/ioutil"
	"os"
	"testing"
)

func TempFile(t testing.TB, dir, pattern string) (f *os.File) {
	t.Helper()
	r0, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func TempDir(t testing.TB, dir, pattern string) (name string) {
	t.Helper()
	r0, err := ioutil.TempDir(dir, pattern)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
