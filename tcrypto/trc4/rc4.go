package trc4

import (
	"crypto/rc4"
	"testing"
)

func NewCipher(t testing.TB, key []byte) *rc4.Cipher {
	t.Helper()
	r0, err := rc4.NewCipher(key)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
