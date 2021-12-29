package tlog

import (
	"log"
	"testing"
)

func Output(t testing.TB, calldepth int, s string) {
	t.Helper()
	err := log.Output(calldepth, s)
	if err != nil {
		t.Fatal(err)
	}
	return
}
