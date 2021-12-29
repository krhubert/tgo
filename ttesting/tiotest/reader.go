package tiotest

import (
	"io"
	"testing"
	"testing/iotest"
)

func TestReader(t testing.TB, r io.Reader, content []byte) {
	t.Helper()
	err := iotest.TestReader(r, content)
	if err != nil {
		t.Fatal(err)
	}
	return
}
