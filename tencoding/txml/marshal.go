package txml

import (
	"encoding/xml"
	"testing"
)

func Marshal(t testing.TB, v any) []byte {
	t.Helper()
	r0, err := xml.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func MarshalIndent(t testing.TB, v any, prefix, indent string) []byte {
	t.Helper()
	r0, err := xml.MarshalIndent(v, prefix, indent)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
