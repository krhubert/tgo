package tquick

import (
	"testing"
	"testing/quick"
)

func Check(t testing.TB, f any, config *quick.Config) {
	t.Helper()
	err := quick.Check(f, config)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func CheckEqual(t testing.TB, f, g any, config *quick.Config) {
	t.Helper()
	err := quick.CheckEqual(f, g, config)
	if err != nil {
		t.Fatal(err)
	}
	return
}
