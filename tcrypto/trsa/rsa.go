package trsa

import (
	"crypto/rsa"
	"hash"
	"io"
	"testing"
)

func GenerateKey(t testing.TB, random io.Reader, bits int) *rsa.PrivateKey {
	t.Helper()
	r0, err := rsa.GenerateKey(random, bits)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func GenerateMultiPrimeKey(t testing.TB, random io.Reader, nprimes int, bits int) *rsa.PrivateKey {
	t.Helper()
	r0, err := rsa.GenerateMultiPrimeKey(random, nprimes, bits)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func EncryptOAEP(t testing.TB, hash hash.Hash, random io.Reader, pub *rsa.PublicKey, msg []byte, label []byte) []byte {
	t.Helper()
	r0, err := rsa.EncryptOAEP(hash, random, pub, msg, label)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DecryptOAEP(t testing.TB, hash hash.Hash, random io.Reader, priv *rsa.PrivateKey, ciphertext []byte, label []byte) []byte {
	t.Helper()
	r0, err := rsa.DecryptOAEP(hash, random, priv, ciphertext, label)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
