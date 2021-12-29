package tasn1

import (
	"encoding/asn1"
	"testing"
)

func Unmarshal(t testing.TB, b []byte, val any) (rest []byte) {
	t.Helper()
	r0, err := asn1.Unmarshal(b, val)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func UnmarshalWithParams(t testing.TB, b []byte, val any, params string) (rest []byte) {
	t.Helper()
	r0, err := asn1.UnmarshalWithParams(b, val, params)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
