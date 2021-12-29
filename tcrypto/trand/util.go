package trand

import (
	"crypto/rand"
	"io"
	"math/big"
	"testing"
)

func Prime(t testing.TB, r io.Reader, bits int) (p *big.Int) {
	t.Helper()
	r0, err := rand.Prime(r, bits)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Int(t testing.TB, r io.Reader, max *big.Int) (n *big.Int) {
	t.Helper()
	r0, err := rand.Int(r, max)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
