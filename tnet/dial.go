package tnet

import (
	"net"
	"testing"
	"time"
)

func Dial(t testing.TB, network, address string) net.Conn {
	t.Helper()
	r0, err := net.Dial(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DialTimeout(t testing.TB, network, address string, timeout time.Duration) net.Conn {
	t.Helper()
	r0, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Listen(t testing.TB, network, address string) net.Listener {
	t.Helper()
	r0, err := net.Listen(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ListenPacket(t testing.TB, network, address string) net.PacketConn {
	t.Helper()
	r0, err := net.ListenPacket(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
