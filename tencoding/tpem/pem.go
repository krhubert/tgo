package tpem

import (
	"encoding/pem"
	"io"
	"testing"
)

func Encode(t testing.TB, out io.Writer, b *pem.Block) {
	t.Helper()
	err := pem.Encode(out, b)
	if err != nil {
		t.Fatal(err)
	}
	return
}
