package tos

import (
	"os"
	"testing"
)

func Getgroups(t testing.TB) []int {
	t.Helper()
	r0, err := os.Getgroups()
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
