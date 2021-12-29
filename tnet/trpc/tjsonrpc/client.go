package tjsonrpc

import (
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

func Dial(t testing.TB, network, address string) *rpc.Client {
	t.Helper()
	r0, err := jsonrpc.Dial(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
