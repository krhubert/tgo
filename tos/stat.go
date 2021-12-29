package tos

import (
	"os"
	"testing"
)

func Stat(t testing.TB, name string) os.FileInfo {
	t.Helper()
	r0, err := os.Stat(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Lstat(t testing.TB, name string) os.FileInfo {
	t.Helper()
	r0, err := os.Lstat(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
