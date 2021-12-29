package tsmtp

import (
	"net"
	"net/smtp"
	"testing"
)

func Dial(t testing.TB, addr string) *smtp.Client {
	t.Helper()
	r0, err := smtp.Dial(addr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func NewClient(t testing.TB, conn net.Conn, host string) *smtp.Client {
	t.Helper()
	r0, err := smtp.NewClient(conn, host)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func SendMail(t testing.TB, addr string, a smtp.Auth, from string, to []string, msg []byte) {
	t.Helper()
	err := smtp.SendMail(addr, a, from, to, msg)
	if err != nil {
		t.Fatal(err)
	}
	return
}
