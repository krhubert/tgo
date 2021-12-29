package thex

import (
	"encoding/hex"
	"testing"
)

func Decode(t testing.TB, dst, src []byte) int {
	t.Helper()
	r0, err := hex.Decode(dst, src)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DecodeString(t testing.TB, s string) []byte {
	t.Helper()
	r0, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
