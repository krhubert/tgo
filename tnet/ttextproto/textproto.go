package ttextproto

import (
	"net/textproto"
	"testing"
)

func Dial(t testing.TB, network, addr string) *textproto.Conn {
	t.Helper()
	r0, err := textproto.Dial(network, addr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
