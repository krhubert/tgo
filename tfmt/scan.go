package tfmt

import (
	"fmt"
	"io"
	"testing"
)

func Scan(t testing.TB, a ...any) (n int) {
	t.Helper()
	r0, err := fmt.Scan(a)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Scanln(t testing.TB, a ...any) (n int) {
	t.Helper()
	r0, err := fmt.Scanln(a)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Scanf(t testing.TB, format string, a ...any) (n int) {
	t.Helper()
	r0, err := fmt.Scanf(format, a)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Sscan(t testing.TB, str string, a ...any) (n int) {
	t.Helper()
	r0, err := fmt.Sscan(str, a)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Sscanln(t testing.TB, str string, a ...any) (n int) {
	t.Helper()
	r0, err := fmt.Sscanln(str, a)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Sscanf(t testing.TB, str string, format string, a ...any) (n int) {
	t.Helper()
	r0, err := fmt.Sscanf(str, format, a)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Fscan(t testing.TB, r io.Reader, a ...any) (n int) {
	t.Helper()
	r0, err := fmt.Fscan(r, a)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Fscanln(t testing.TB, r io.Reader, a ...any) (n int) {
	t.Helper()
	r0, err := fmt.Fscanln(r, a)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Fscanf(t testing.TB, r io.Reader, format string, a ...any) (n int) {
	t.Helper()
	r0, err := fmt.Fscanf(r, format, a)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
