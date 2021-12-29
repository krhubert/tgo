package tx509

import (
	"crypto/x509"
	"testing"
)

func SystemCertPool(t testing.TB) *x509.CertPool {
	t.Helper()
	r0, err := x509.SystemCertPool()
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
