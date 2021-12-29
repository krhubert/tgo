package taes

import (
	"crypto/aes"
	"crypto/cipher"
	"testing"
)

func NewCipher(t testing.TB, key []byte) cipher.Block {
	t.Helper()
	r0, err := aes.NewCipher(key)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
