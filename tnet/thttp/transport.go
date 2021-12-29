package thttp

import (
	"net/http"
	"net/url"
	"testing"
)

func ProxyFromEnvironment(t testing.TB, req *http.Request) *url.URL {
	t.Helper()
	r0, err := http.ProxyFromEnvironment(req)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
