package tx509

import (
	"crypto/x509"
	"encoding/pem"
	"io"
	"testing"
)

func DecryptPEMBlock(t testing.TB, b *pem.Block, password []byte) []byte {
	t.Helper()
	r0, err := x509.DecryptPEMBlock(b, password)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func EncryptPEMBlock(t testing.TB, rand io.Reader, blockType string, data, password []byte, alg x509.PEMCipher) *pem.Block {
	t.Helper()
	r0, err := x509.EncryptPEMBlock(rand, blockType, data, password, alg)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
