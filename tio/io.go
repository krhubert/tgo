package tio

import (
	"io"
	"testing"
)

func WriteString(t testing.TB, w io.Writer, s string) (n int) {
	t.Helper()
	r0, err := io.WriteString(w, s)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ReadAtLeast(t testing.TB, r io.Reader, buf []byte, min int) (n int) {
	t.Helper()
	r0, err := io.ReadAtLeast(r, buf, min)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ReadFull(t testing.TB, r io.Reader, buf []byte) (n int) {
	t.Helper()
	r0, err := io.ReadFull(r, buf)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func CopyN(t testing.TB, dst io.Writer, src io.Reader, n int64) (written int64) {
	t.Helper()
	r0, err := io.CopyN(dst, src, n)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Copy(t testing.TB, dst io.Writer, src io.Reader) (written int64) {
	t.Helper()
	r0, err := io.Copy(dst, src)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func CopyBuffer(t testing.TB, dst io.Writer, src io.Reader, buf []byte) (written int64) {
	t.Helper()
	r0, err := io.CopyBuffer(dst, src, buf)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ReadAll(t testing.TB, r io.Reader) []byte {
	t.Helper()
	r0, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
