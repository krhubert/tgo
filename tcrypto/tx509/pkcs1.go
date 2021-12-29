package tx509

import (
	"crypto/rsa"
	"crypto/x509"
	"testing"
)

func ParsePKCS1PrivateKey(t testing.TB, der []byte) *rsa.PrivateKey {
	t.Helper()
	r0, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParsePKCS1PublicKey(t testing.TB, der []byte) *rsa.PublicKey {
	t.Helper()
	r0, err := x509.ParsePKCS1PublicKey(der)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
