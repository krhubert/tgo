package tbuildinfo

import (
	"debug/buildinfo"
	"io"
	"testing"
)

func ReadFile(t testing.TB, name string) (info *buildinfo.BuildInfo) {
	t.Helper()
	r0, err := buildinfo.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func Read(t testing.TB, r io.ReaderAt) *buildinfo.BuildInfo {
	t.Helper()
	r0, err := buildinfo.Read(r)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
