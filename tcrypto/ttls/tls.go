package ttls

import (
	"crypto/tls"
	"net"
	"testing"
)

func Listen(t testing.TB, network, laddr string, config *tls.Config) net.Listener {
	t.Helper()
	r0, err := tls.Listen(network, laddr, config)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func DialWithDialer(t testing.TB, dialer *net.Dialer, network, addr string, config *tls.Config) *tls.Conn {
	t.Helper()
	r0, err := tls.DialWithDialer(dialer, network, addr, config)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Dial(t testing.TB, network, addr string, config *tls.Config) *tls.Conn {
	t.Helper()
	r0, err := tls.Dial(network, addr, config)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func LoadX509KeyPair(t testing.TB, certFile, keyFile string) tls.Certificate {
	t.Helper()
	r0, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func X509KeyPair(t testing.TB, certPEMBlock, keyPEMBlock []byte) tls.Certificate {
	t.Helper()
	r0, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
