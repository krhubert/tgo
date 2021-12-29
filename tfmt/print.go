package tfmt

import (
	"fmt"
	"io"
	"testing"
)

func Fprintf(t testing.TB, w io.Writer, format string, a ...any) (n int) {
	t.Helper()
	r0, err := fmt.Fprintf(w, format, a)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Printf(t testing.TB, format string, a ...any) (n int) {
	t.Helper()
	r0, err := fmt.Printf(format, a)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Fprint(t testing.TB, w io.Writer, a ...any) (n int) {
	t.Helper()
	r0, err := fmt.Fprint(w, a)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Print(t testing.TB, a ...any) (n int) {
	t.Helper()
	r0, err := fmt.Print(a)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Fprintln(t testing.TB, w io.Writer, a ...any) (n int) {
	t.Helper()
	r0, err := fmt.Fprintln(w, a)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Println(t testing.TB, a ...any) (n int) {
	t.Helper()
	r0, err := fmt.Println(a)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
