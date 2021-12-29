package tbinary

import (
	"encoding/binary"
	"io"
	"testing"
)

func ReadUvarint(t testing.TB, r io.ByteReader) uint64 {
	t.Helper()
	r0, err := binary.ReadUvarint(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ReadVarint(t testing.TB, r io.ByteReader) int64 {
	t.Helper()
	r0, err := binary.ReadVarint(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
