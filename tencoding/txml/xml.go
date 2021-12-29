package txml

import (
	"encoding/xml"
	"io"
	"testing"
)

func EscapeText(t testing.TB, w io.Writer, s []byte) {
	t.Helper()
	err := xml.EscapeText(w, s)
	if err != nil {
		t.Fatal(err)
	}
	return
}
