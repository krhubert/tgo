package tdsa

import (
	"crypto/dsa"
	"io"
	"math/big"
	"testing"
)

func GenerateParameters(t testing.TB, params *dsa.Parameters, rand io.Reader, sizes dsa.ParameterSizes) {
	t.Helper()
	err := dsa.GenerateParameters(params, rand, sizes)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func GenerateKey(t testing.TB, priv *dsa.PrivateKey, rand io.Reader) {
	t.Helper()
	err := dsa.GenerateKey(priv, rand)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func Sign(t testing.TB, rand io.Reader, priv *dsa.PrivateKey, hash []byte) (r, s *big.Int) {
	t.Helper()
	r0, r1, err := dsa.Sign(rand, priv, hash)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1
}
