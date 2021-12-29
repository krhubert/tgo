package trsa

import (
	"crypto"
	"crypto/rsa"
	"io"
	"testing"
)

func EncryptPKCS1v15(t testing.TB, rand io.Reader, pub *rsa.PublicKey, msg []byte) []byte {
	t.Helper()
	r0, err := rsa.EncryptPKCS1v15(rand, pub, msg)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DecryptPKCS1v15(t testing.TB, rand io.Reader, priv *rsa.PrivateKey, ciphertext []byte) []byte {
	t.Helper()
	r0, err := rsa.DecryptPKCS1v15(rand, priv, ciphertext)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DecryptPKCS1v15SessionKey(t testing.TB, rand io.Reader, priv *rsa.PrivateKey, ciphertext []byte, key []byte) {
	t.Helper()
	err := rsa.DecryptPKCS1v15SessionKey(rand, priv, ciphertext, key)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func SignPKCS1v15(t testing.TB, rand io.Reader, priv *rsa.PrivateKey, hash crypto.Hash, hashed []byte) []byte {
	t.Helper()
	r0, err := rsa.SignPKCS1v15(rand, priv, hash, hashed)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func VerifyPKCS1v15(t testing.TB, pub *rsa.PublicKey, hash crypto.Hash, hashed []byte, sig []byte) {
	t.Helper()
	err := rsa.VerifyPKCS1v15(pub, hash, hashed, sig)
	if err != nil {
		t.Fatal(err)
	}
	return
}
