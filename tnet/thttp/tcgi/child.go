package tcgi

import (
	"net/http"
	"net/http/cgi"
	"testing"
)

func Request(t testing.TB) *http.Request {
	t.Helper()
	r0, err := cgi.Request()
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func RequestFromMap(t testing.TB, params map[string]string) *http.Request {
	t.Helper()
	r0, err := cgi.RequestFromMap(params)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Serve(t testing.TB, handler http.Handler) {
	t.Helper()
	err := cgi.Serve(handler)
	if err != nil {
		t.Fatal(err)
	}
	return
}
