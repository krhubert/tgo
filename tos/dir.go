package tos

import (
	"os"
	"testing"
)

func ReadDir(t testing.TB, name string) []os.DirEntry {
	t.Helper()
	r0, err := os.ReadDir(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
