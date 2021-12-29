package trpc

import (
	"net/rpc"
	"testing"
)

func Register(t testing.TB, rcvr any) {
	t.Helper()
	err := rpc.Register(rcvr)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func RegisterName(t testing.TB, name string, rcvr any) {
	t.Helper()
	err := rpc.RegisterName(name, rcvr)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func ServeRequest(t testing.TB, codec rpc.ServerCodec) {
	t.Helper()
	err := rpc.ServeRequest(codec)
	if err != nil {
		t.Fatal(err)
	}
	return
}
