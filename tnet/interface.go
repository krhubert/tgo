package tnet

import (
	"net"
	"testing"
)

func Interfaces(t testing.TB) []net.Interface {
	t.Helper()
	r0, err := net.Interfaces()
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func InterfaceAddrs(t testing.TB) []net.Addr {
	t.Helper()
	r0, err := net.InterfaceAddrs()
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func InterfaceByIndex(t testing.TB, index int) *net.Interface {
	t.Helper()
	r0, err := net.InterfaceByIndex(index)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func InterfaceByName(t testing.TB, name string) *net.Interface {
	t.Helper()
	r0, err := net.InterfaceByName(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
