package tasn1

import (
	"encoding/asn1"
	"testing"
)

func Marshal(t testing.TB, val any) []byte {
	t.Helper()
	r0, err := asn1.Marshal(val)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func MarshalWithParams(t testing.TB, val any, params string) []byte {
	t.Helper()
	r0, err := asn1.MarshalWithParams(val, params)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
