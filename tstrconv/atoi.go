package tstrconv

import (
	"strconv"
	"testing"
)

func ParseUint(t testing.TB, s string, base int, bitSize int) uint64 {
	t.Helper()
	r0, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseInt(t testing.TB, s string, base int, bitSize int) (i int64) {
	t.Helper()
	r0, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Atoi(t testing.TB, s string) int {
	t.Helper()
	r0, err := strconv.Atoi(s)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
