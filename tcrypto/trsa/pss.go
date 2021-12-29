package trsa

import (
	"crypto"
	"crypto/rsa"
	"io"
	"testing"
)

func SignPSS(t testing.TB, rand io.Reader, priv *rsa.PrivateKey, hash crypto.Hash, digest []byte, opts *rsa.PSSOptions) []byte {
	t.Helper()
	r0, err := rsa.SignPSS(rand, priv, hash, digest, opts)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func VerifyPSS(t testing.TB, pub *rsa.PublicKey, hash crypto.Hash, digest []byte, sig []byte, opts *rsa.PSSOptions) {
	t.Helper()
	err := rsa.VerifyPSS(pub, hash, digest, sig, opts)
	if err != nil {
		t.Fatal(err)
	}
	return
}
