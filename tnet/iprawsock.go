package tnet

import (
	"net"
	"testing"
)

func ResolveIPAddr(t testing.TB, network, address string) *net.IPAddr {
	t.Helper()
	r0, err := net.ResolveIPAddr(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DialIP(t testing.TB, network string, laddr, raddr *net.IPAddr) *net.IPConn {
	t.Helper()
	r0, err := net.DialIP(network, laddr, raddr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ListenIP(t testing.TB, network string, laddr *net.IPAddr) *net.IPConn {
	t.Helper()
	r0, err := net.ListenIP(network, laddr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
