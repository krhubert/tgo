package truntime

import (
	"runtime"
	"testing"
)

func StartTrace(t testing.TB) {
	t.Helper()
	err := runtime.StartTrace()
	if err != nil {
		t.Fatal(err)
	}
	return
}
