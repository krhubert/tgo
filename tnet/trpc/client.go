package trpc

import (
	"net/rpc"
	"testing"
)

func DialHTTP(t testing.TB, network, address string) *rpc.Client {
	t.Helper()
	r0, err := rpc.DialHTTP(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DialHTTPPath(t testing.TB, network, address, path string) *rpc.Client {
	t.Helper()
	r0, err := rpc.DialHTTPPath(network, address, path)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Dial(t testing.TB, network, address string) *rpc.Client {
	t.Helper()
	r0, err := rpc.Dial(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
