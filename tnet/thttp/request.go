package thttp

import (
	"bufio"
	"context"
	"io"
	"net/http"
	"testing"
)

func NewRequest(t testing.TB, method, url string, body io.Reader) *http.Request {
	t.Helper()
	r0, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func NewRequestWithContext(t testing.TB, ctx context.Context, method, url string, body io.Reader) *http.Request {
	t.Helper()
	r0, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ReadRequest(t testing.TB, b *bufio.Reader) *http.Request {
	t.Helper()
	r0, err := http.ReadRequest(b)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
