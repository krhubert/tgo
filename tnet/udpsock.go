package tnet

import (
	"net"
	"testing"
)

func ResolveUDPAddr(t testing.TB, network, address string) *net.UDPAddr {
	t.Helper()
	r0, err := net.ResolveUDPAddr(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DialUDP(t testing.TB, network string, laddr, raddr *net.UDPAddr) *net.UDPConn {
	t.Helper()
	r0, err := net.DialUDP(network, laddr, raddr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ListenUDP(t testing.TB, network string, laddr *net.UDPAddr) *net.UDPConn {
	t.Helper()
	r0, err := net.ListenUDP(network, laddr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ListenMulticastUDP(t testing.TB, network string, ifi *net.Interface, gaddr *net.UDPAddr) *net.UDPConn {
	t.Helper()
	r0, err := net.ListenMulticastUDP(network, ifi, gaddr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
