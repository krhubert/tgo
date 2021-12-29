package tx509

import (
	"crypto/x509"
	"testing"
)

func ParsePKCS8PrivateKey(t testing.TB, der []byte) (key any) {
	t.Helper()
	r0, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func MarshalPKCS8PrivateKey(t testing.TB, key any) []byte {
	t.Helper()
	r0, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
