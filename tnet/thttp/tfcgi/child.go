package tfcgi

import (
	"net"
	"net/http"
	"net/http/fcgi"
	"testing"
)

func Serve(t testing.TB, l net.Listener, handler http.Handler) {
	t.Helper()
	err := fcgi.Serve(l, handler)
	if err != nil {
		t.Fatal(err)
	}
	return
}
