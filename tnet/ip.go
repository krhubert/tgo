package tnet

import (
	"net"
	"testing"
)

func ParseCIDR(t testing.TB, s string) (net.IP, *net.IPNet) {
	t.Helper()
	r0, r1, err := net.ParseCIDR(s)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1
}
