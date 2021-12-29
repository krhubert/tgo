package telliptic

import (
	"crypto/elliptic"
	"io"
	"math/big"
	"testing"
)

func GenerateKey(t testing.TB, curve elliptic.Curve, rand io.Reader) (priv []byte, x, y *big.Int) {
	t.Helper()
	r0, r1, r2, err := elliptic.GenerateKey(curve, rand)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1, r2
}
