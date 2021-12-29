package tjson

import (
	"bytes"
	"encoding/json"
	"testing"
)

func Compact(t testing.TB, dst *bytes.Buffer, src []byte) {
	t.Helper()
	err := json.Compact(dst, src)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func Indent(t testing.TB, dst *bytes.Buffer, src []byte, prefix, indent string) {
	t.Helper()
	err := json.Indent(dst, src, prefix, indent)
	if err != nil {
		t.Fatal(err)
	}
	return
}
