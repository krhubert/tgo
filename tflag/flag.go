package tflag

import (
	"flag"
	"testing"
)

func Set(t testing.TB, name, value string) {
	t.Helper()
	err := flag.Set(name, value)
	if err != nil {
		t.Fatal(err)
	}
	return
}
