package tnet

import (
	"net"
	"testing"
)

func ResolveUnixAddr(t testing.TB, network, address string) *net.UnixAddr {
	t.Helper()
	r0, err := net.ResolveUnixAddr(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DialUnix(t testing.TB, network string, laddr, raddr *net.UnixAddr) *net.UnixConn {
	t.Helper()
	r0, err := net.DialUnix(network, laddr, raddr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ListenUnix(t testing.TB, network string, laddr *net.UnixAddr) *net.UnixListener {
	t.Helper()
	r0, err := net.ListenUnix(network, laddr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ListenUnixgram(t testing.TB, network string, laddr *net.UnixAddr) *net.UnixConn {
	t.Helper()
	r0, err := net.ListenUnixgram(network, laddr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
