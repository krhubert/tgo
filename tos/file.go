package tos

import (
	"os"
	"testing"
)

func Mkdir(t testing.TB, name string, perm os.FileMode) {
	t.Helper()
	err := os.Mkdir(name, perm)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func Chdir(t testing.TB, dir string) {
	t.Helper()
	err := os.Chdir(dir)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func Open(t testing.TB, name string) *os.File {
	t.Helper()
	r0, err := os.Open(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Create(t testing.TB, name string) *os.File {
	t.Helper()
	r0, err := os.Create(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func OpenFile(t testing.TB, name string, flag int, perm os.FileMode) *os.File {
	t.Helper()
	r0, err := os.OpenFile(name, flag, perm)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Rename(t testing.TB, oldpath, newpath string) {
	t.Helper()
	err := os.Rename(oldpath, newpath)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func UserCacheDir(t testing.TB) string {
	t.Helper()
	r0, err := os.UserCacheDir()
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func UserConfigDir(t testing.TB) string {
	t.Helper()
	r0, err := os.UserConfigDir()
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func UserHomeDir(t testing.TB) string {
	t.Helper()
	r0, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Chmod(t testing.TB, name string, mode os.FileMode) {
	t.Helper()
	err := os.Chmod(name, mode)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func ReadFile(t testing.TB, name string) []byte {
	t.Helper()
	r0, err := os.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func WriteFile(t testing.TB, name string, data []byte, perm os.FileMode) {
	t.Helper()
	err := os.WriteFile(name, data, perm)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func Pipe(t testing.TB) (r *os.File, w *os.File) {
	t.Helper()
	r0, r1, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	return r0, r1
}
