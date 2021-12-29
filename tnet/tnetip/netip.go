package tnetip

import (
	"net/netip"
	"testing"
)

func ParseAddr(t testing.TB, s string) netip.Addr {
	t.Helper()
	r0, err := netip.ParseAddr(s)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseAddrPort(t testing.TB, s string) netip.AddrPort {
	t.Helper()
	r0, err := netip.ParseAddrPort(s)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParsePrefix(t testing.TB, s string) netip.Prefix {
	t.Helper()
	r0, err := netip.ParsePrefix(s)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
