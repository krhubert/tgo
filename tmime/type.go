package tmime

import (
	"mime"
	"testing"
)

func ExtensionsByType(t testing.TB, typ string) []string {
	t.Helper()
	r0, err := mime.ExtensionsByType(typ)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func AddExtensionType(t testing.TB, ext, typ string) {
	t.Helper()
	err := mime.AddExtensionType(ext, typ)
	if err != nil {
		t.Fatal(err)
	}
	return
}
