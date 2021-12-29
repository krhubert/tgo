package tstrconv

import (
	"strconv"
	"testing"
)

func ParseComplex(t testing.TB, s string, bitSize int) complex128 {
	t.Helper()
	r0, err := strconv.ParseComplex(s, bitSize)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
