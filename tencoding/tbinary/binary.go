package tbinary

import (
	"encoding/binary"
	"io"
	"testing"
)

func Read(t testing.TB, r io.Reader, order binary.ByteOrder, data any) {
	t.Helper()
	err := binary.Read(r, order, data)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func Write(t testing.TB, w io.Writer, order binary.ByteOrder, data any) {
	t.Helper()
	err := binary.Write(w, order, data)
	if err != nil {
		t.Fatal(err)
	}
	return
}
