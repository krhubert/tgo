package tmime

import (
	"mime"
	"testing"
)

func ParseMediaType(t testing.TB, v string) (mediatype string, params map[string]string) {
	t.Helper()
	r0, r1, err := mime.ParseMediaType(v)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1
}
