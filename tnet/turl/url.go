package turl

import (
	"net/url"
	"testing"
)

func QueryUnescape(t testing.TB, s string) string {
	t.Helper()
	r0, err := url.QueryUnescape(s)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func PathUnescape(t testing.TB, s string) string {
	t.Helper()
	r0, err := url.PathUnescape(s)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Parse(t testing.TB, rawURL string) *url.URL {
	t.Helper()
	r0, err := url.Parse(rawURL)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseRequestURI(t testing.TB, rawURL string) *url.URL {
	t.Helper()
	r0, err := url.ParseRequestURI(rawURL)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseQuery(t testing.TB, query string) url.Values {
	t.Helper()
	r0, err := url.ParseQuery(query)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
