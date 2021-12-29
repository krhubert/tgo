package thttp

import (
	"io"
	"net/http"
	"net/url"
	"testing"
)

func Get(t testing.TB, url string) (resp *http.Response) {
	t.Helper()
	r0, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Post(t testing.TB, url, contentType string, body io.Reader) (resp *http.Response) {
	t.Helper()
	r0, err := http.Post(url, contentType, body)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func PostForm(t testing.TB, url string, data url.Values) (resp *http.Response) {
	t.Helper()
	r0, err := http.PostForm(url, data)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Head(t testing.TB, url string) (resp *http.Response) {
	t.Helper()
	r0, err := http.Head(url)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
