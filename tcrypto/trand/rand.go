package trand

import (
	"crypto/rand"
	"testing"
)

func Read(t testing.TB, b []byte) (n int) {
	t.Helper()
	r0, err := rand.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
