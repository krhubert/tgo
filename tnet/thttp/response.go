package thttp

import (
	"bufio"
	"net/http"
	"testing"
)

func ReadResponse(t testing.TB, r *bufio.Reader, req *http.Request) *http.Response {
	t.Helper()
	r0, err := http.ReadResponse(r, req)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
