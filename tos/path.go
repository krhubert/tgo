package tos

import (
	"os"
	"testing"
)

func MkdirAll(t testing.TB, path string, perm os.FileMode) {
	t.Helper()
	err := os.MkdirAll(path, perm)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func RemoveAll(t testing.TB, path string) {
	t.Helper()
	err := os.RemoveAll(path)
	if err != nil {
		t.Fatal(err)
	}
	return
}
