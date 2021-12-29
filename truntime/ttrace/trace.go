package ttrace

import (
	"io"
	"runtime/trace"
	"testing"
)

func Start(t testing.TB, w io.Writer) {
	t.Helper()
	err := trace.Start(w)
	if err != nil {
		t.Fatal(err)
	}
	return
}
