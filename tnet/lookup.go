package tnet

import (
	"net"
	"testing"
)

func LookupHost(t testing.TB, host string) (addrs []string) {
	t.Helper()
	r0, err := net.LookupHost(host)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func LookupIP(t testing.TB, host string) []net.IP {
	t.Helper()
	r0, err := net.LookupIP(host)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func LookupPort(t testing.TB, network, service string) (port int) {
	t.Helper()
	r0, err := net.LookupPort(network, service)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func LookupCNAME(t testing.TB, host string) (cname string) {
	t.Helper()
	r0, err := net.LookupCNAME(host)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func LookupSRV(t testing.TB, service, proto, name string) (cname string, addrs []*net.SRV) {
	t.Helper()
	r0, r1, err := net.LookupSRV(service, proto, name)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1
}

func LookupMX(t testing.TB, name string) []*net.MX {
	t.Helper()
	r0, err := net.LookupMX(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func LookupNS(t testing.TB, name string) []*net.NS {
	t.Helper()
	r0, err := net.LookupNS(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func LookupTXT(t testing.TB, name string) []string {
	t.Helper()
	r0, err := net.LookupTXT(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func LookupAddr(t testing.TB, addr string) (names []string) {
	t.Helper()
	r0, err := net.LookupAddr(addr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
