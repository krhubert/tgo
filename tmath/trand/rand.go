package trand

import (
	"math/rand"
	"testing"
)

func Read(t testing.TB, p []byte) (n int) {
	t.Helper()
	r0, err := rand.Read(p)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
