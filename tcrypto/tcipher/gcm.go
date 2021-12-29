package tcipher

import (
	"crypto/cipher"
	"testing"
)

func NewGCM(t testing.TB, block cipher.Block) cipher.AEAD {
	t.Helper()
	r0, err := cipher.NewGCM(block)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func NewGCMWithNonceSize(t testing.TB, block cipher.Block, size int) cipher.AEAD {
	t.Helper()
	r0, err := cipher.NewGCMWithNonceSize(block, size)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func NewGCMWithTagSize(t testing.TB, block cipher.Block, tagSize int) cipher.AEAD {
	t.Helper()
	r0, err := cipher.NewGCMWithTagSize(block, tagSize)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
