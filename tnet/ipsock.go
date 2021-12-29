package tnet

import (
	"net"
	"testing"
)

func SplitHostPort(t testing.TB, hostport string) (host, port string) {
	t.Helper()
	r0, r1, err := net.SplitHostPort(hostport)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1
}
