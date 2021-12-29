package tos

import (
	"os"
	"testing"
)

func Executable(t testing.TB) string {
	t.Helper()
	r0, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
