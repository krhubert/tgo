package tnet

import (
	"net"
	"testing"
)

func ResolveTCPAddr(t testing.TB, network, address string) *net.TCPAddr {
	t.Helper()
	r0, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DialTCP(t testing.TB, network string, laddr, raddr *net.TCPAddr) *net.TCPConn {
	t.Helper()
	r0, err := net.DialTCP(network, laddr, raddr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ListenTCP(t testing.TB, network string, laddr *net.TCPAddr) *net.TCPListener {
	t.Helper()
	r0, err := net.ListenTCP(network, laddr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
