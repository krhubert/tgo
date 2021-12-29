package tos

import (
	"os"
	"testing"
)

func CreateTemp(t testing.TB, dir, pattern string) *os.File {
	t.Helper()
	r0, err := os.CreateTemp(dir, pattern)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func MkdirTemp(t testing.TB, dir, pattern string) string {
	t.Helper()
	r0, err := os.MkdirTemp(dir, pattern)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
