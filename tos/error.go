package tos

import (
	"os"
	"testing"
)

func NewSyscallError(t testing.TB, syscall string, err error) {
	t.Helper()
	err1 := os.NewSyscallError(syscall, err)
	if err1 != nil {
		t.Fatal(err1)
	}
	return
}
