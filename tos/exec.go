package tos

import (
	"os"
	"testing"
)

func FindProcess(t testing.TB, pid int) *os.Process {
	t.Helper()
	r0, err := os.FindProcess(pid)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func StartProcess(t testing.TB, name string, argv []string, attr *os.ProcAttr) *os.Process {
	t.Helper()
	r0, err := os.StartProcess(name, argv, attr)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
