package thttp

import (
	"net"
	"net/http"
	"testing"
)

func Serve(t testing.TB, l net.Listener, handler http.Handler) {
	t.Helper()
	err := http.Serve(l, handler)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func ServeTLS(t testing.TB, l net.Listener, handler http.Handler, certFile, keyFile string) {
	t.Helper()
	err := http.ServeTLS(l, handler, certFile, keyFile)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func ListenAndServe(t testing.TB, addr string, handler http.Handler) {
	t.Helper()
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func ListenAndServeTLS(t testing.TB, addr, certFile, keyFile string, handler http.Handler) {
	t.Helper()
	err := http.ListenAndServeTLS(addr, certFile, keyFile, handler)
	if err != nil {
		t.Fatal(err)
	}
	return
}
