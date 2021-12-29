package tx509

import (
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"io"
	"testing"
)

func ParsePKIXPublicKey(t testing.TB, derBytes []byte) (pub any) {
	t.Helper()
	r0, err := x509.ParsePKIXPublicKey(derBytes)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func MarshalPKIXPublicKey(t testing.TB, pub any) []byte {
	t.Helper()
	r0, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func CreateCertificate(t testing.TB, rand io.Reader, template, parent *x509.Certificate, pub, priv any) []byte {
	t.Helper()
	r0, err := x509.CreateCertificate(rand, template, parent, pub, priv)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseCRL(t testing.TB, crlBytes []byte) *pkix.CertificateList {
	t.Helper()
	r0, err := x509.ParseCRL(crlBytes)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseDERCRL(t testing.TB, derBytes []byte) *pkix.CertificateList {
	t.Helper()
	r0, err := x509.ParseDERCRL(derBytes)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func CreateCertificateRequest(t testing.TB, rand io.Reader, template *x509.CertificateRequest, priv any) (csr []byte) {
	t.Helper()
	r0, err := x509.CreateCertificateRequest(rand, template, priv)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseCertificateRequest(t testing.TB, asn1Data []byte) *x509.CertificateRequest {
	t.Helper()
	r0, err := x509.ParseCertificateRequest(asn1Data)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func CreateRevocationList(t testing.TB, rand io.Reader, template *x509.RevocationList, issuer *x509.Certificate, priv crypto.Signer) []byte {
	t.Helper()
	r0, err := x509.CreateRevocationList(rand, template, issuer, priv)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
