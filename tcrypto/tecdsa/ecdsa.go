package tecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"io"
	"math/big"
	"testing"
)

func GenerateKey(t testing.TB, c elliptic.Curve, rand io.Reader) *ecdsa.PrivateKey {
	t.Helper()
	r0, err := ecdsa.GenerateKey(c, rand)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Sign(t testing.TB, rand io.Reader, priv *ecdsa.PrivateKey, hash []byte) (r, s *big.Int) {
	t.Helper()
	r0, r1, err := ecdsa.Sign(rand, priv, hash)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1
}

func SignASN1(t testing.TB, rand io.Reader, priv *ecdsa.PrivateKey, hash []byte) []byte {
	t.Helper()
	r0, err := ecdsa.SignASN1(rand, priv, hash)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
