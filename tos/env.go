package tos

import (
	"os"
	"testing"
)

func Setenv(t testing.TB, key, value string) {
	t.Helper()
	err := os.Setenv(key, value)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func Unsetenv(t testing.TB, key string) {
	t.Helper()
	err := os.Unsetenv(key)
	if err != nil {
		t.Fatal(err)
	}
	return
}
