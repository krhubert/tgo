package tbufio

import (
	"bufio"
	"testing"
)

func ScanBytes(t testing.TB, data []byte, atEOF bool) (advance int, token []byte) {
	t.Helper()
	r0, r1, err := bufio.ScanBytes(data, atEOF)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1
}

func ScanRunes(t testing.TB, data []byte, atEOF bool) (advance int, token []byte) {
	t.Helper()
	r0, r1, err := bufio.ScanRunes(data, atEOF)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1
}

func ScanLines(t testing.TB, data []byte, atEOF bool) (advance int, token []byte) {
	t.Helper()
	r0, r1, err := bufio.ScanLines(data, atEOF)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1
}

func ScanWords(t testing.TB, data []byte, atEOF bool) (advance int, token []byte) {
	t.Helper()
	r0, r1, err := bufio.ScanWords(data, atEOF)
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1
}
