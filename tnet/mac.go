package tnet

import (
	"net"
	"testing"
)

func ParseMAC(t testing.TB, s string) (hw net.HardwareAddr) {
	t.Helper()
	r0, err := net.ParseMAC(s)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
