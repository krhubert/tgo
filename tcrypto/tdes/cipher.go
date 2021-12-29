package tdes

import (
	"crypto/cipher"
	"crypto/des"
	"testing"
)

func NewCipher(t testing.TB, key []byte) cipher.Block {
	t.Helper()
	r0, err := des.NewCipher(key)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func NewTripleDESCipher(t testing.TB, key []byte) cipher.Block {
	t.Helper()
	r0, err := des.NewTripleDESCipher(key)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
