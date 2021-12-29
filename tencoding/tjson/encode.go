package tjson

import (
	"encoding/json"
	"testing"
)

func Marshal(t testing.TB, v any) []byte {
	t.Helper()
	r0, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func MarshalIndent(t testing.TB, v any, prefix, indent string) []byte {
	t.Helper()
	r0, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
