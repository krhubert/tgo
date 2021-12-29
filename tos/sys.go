package tos

import (
	"os"
	"testing"
)

func Hostname(t testing.TB) (name string) {
	t.Helper()
	r0, err := os.Hostname()
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
