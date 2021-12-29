package tos

import (
	"os"
	"testing"
)

func Getwd(t testing.TB) (dir string) {
	t.Helper()
	r0, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
