package tx509

import (
	"crypto/x509"
	"testing"
)

func ParseCertificate(t testing.TB, der []byte) *x509.Certificate {
	t.Helper()
	r0, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseCertificates(t testing.TB, der []byte) []*x509.Certificate {
	t.Helper()
	r0, err := x509.ParseCertificates(der)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
