package tmail

import (
	"io"
	"net/mail"
	"testing"
	"time"
)

func ReadMessage(t testing.TB, r io.Reader) (msg *mail.Message) {
	t.Helper()
	r0, err := mail.ReadMessage(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseDate(t testing.TB, date string) time.Time {
	t.Helper()
	r0, err := mail.ParseDate(date)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseAddress(t testing.TB, address string) *mail.Address {
	t.Helper()
	r0, err := mail.ParseAddress(address)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseAddressList(t testing.TB, list string) []*mail.Address {
	t.Helper()
	r0, err := mail.ParseAddressList(list)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
