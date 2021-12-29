package tstrconv

import (
	"strconv"
	"testing"
)

func ParseFloat(t testing.TB, s string, bitSize int) float64 {
	t.Helper()
	r0, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
