package tnet

import (
	"net"
	"os"
	"testing"
)

func FileConn(t testing.TB, f *os.File) (c net.Conn) {
	t.Helper()
	r0, err := net.FileConn(f)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func FileListener(t testing.TB, f *os.File) (ln net.Listener) {
	t.Helper()
	r0, err := net.FileListener(f)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func FilePacketConn(t testing.TB, f *os.File) (c net.PacketConn) {
	t.Helper()
	r0, err := net.FilePacketConn(f)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
