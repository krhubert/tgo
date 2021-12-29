package thttp

import (
	"net/http"
	"testing"
	"time"
)

func ParseTime(t testing.TB, text string) time.Time {
	t.Helper()
	r0, err := http.ParseTime(text)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
