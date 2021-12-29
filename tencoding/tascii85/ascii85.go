package tascii85

import (
	"encoding/ascii85"
	"testing"
)

func Decode(t testing.TB, dst, src []byte, flush bool) (ndst, nsrc int) {
	t.Helper()
	r0, r1, err := ascii85.Decode(dst, src, flush)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1
}
