package tx509

import (
	"crypto/ecdsa"
	"crypto/x509"
	"testing"
)

func ParseECPrivateKey(t testing.TB, der []byte) *ecdsa.PrivateKey {
	t.Helper()
	r0, err := x509.ParseECPrivateKey(der)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func MarshalECPrivateKey(t testing.TB, key *ecdsa.PrivateKey) []byte {
	t.Helper()
	r0, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
