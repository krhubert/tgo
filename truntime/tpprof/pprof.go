package tpprof

import (
	"io"
	"runtime/pprof"
	"testing"
)

func WriteHeapProfile(t testing.TB, w io.Writer) {
	t.Helper()
	err := pprof.WriteHeapProfile(w)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func StartCPUProfile(t testing.TB, w io.Writer) {
	t.Helper()
	err := pprof.StartCPUProfile(w)
	if err != nil {
		t.Fatal(err)
	}
	return
}
