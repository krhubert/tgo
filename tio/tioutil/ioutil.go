package tioutil

import (
	"io"
	"io/fs"
	"io/ioutil"
	"testing"
)

func ReadAll(t testing.TB, r io.Reader) []byte {
	t.Helper()
	r0, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ReadFile(t testing.TB, filename string) []byte {
	t.Helper()
	r0, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func WriteFile(t testing.TB, filename string, data []byte, perm fs.FileMode) {
	t.Helper()
	err := ioutil.WriteFile(filename, data, perm)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func ReadDir(t testing.TB, dirname string) []fs.FileInfo {
	t.Helper()
	r0, err := ioutil.ReadDir(dirname)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
