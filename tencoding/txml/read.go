package txml

import (
	"encoding/xml"
	"testing"
)

func Unmarshal(t testing.TB, data []byte, v any) {
	t.Helper()
	err := xml.Unmarshal(data, v)
	if err != nil {
		t.Fatal(err)
	}
	return
}
