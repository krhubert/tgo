package thttputil

import (
	"net/http"
	"net/http/httputil"
	"testing"
)

func DumpRequestOut(t testing.TB, req *http.Request, body bool) []byte {
	t.Helper()
	r0, err := httputil.DumpRequestOut(req, body)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DumpRequest(t testing.TB, req *http.Request, body bool) []byte {
	t.Helper()
	r0, err := httputil.DumpRequest(req, body)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DumpResponse(t testing.TB, resp *http.Response, body bool) []byte {
	t.Helper()
	r0, err := httputil.DumpResponse(resp, body)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
