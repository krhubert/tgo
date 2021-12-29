package ted25519

import (
	"crypto/ed25519"
	"io"
	"testing"
)

func GenerateKey(t testing.TB, rand io.Reader) (ed25519.PublicKey, ed25519.PrivateKey) {
	t.Helper()
	r0, r1, err := ed25519.GenerateKey(rand)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1
}
