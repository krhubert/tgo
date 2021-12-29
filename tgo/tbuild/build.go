package tbuild

import (
	"go/build"
	"testing"
)

func Import(t testing.TB, path, srcDir string, mode build.ImportMode) *build.Package {
	t.Helper()
	r0, err := build.Import(path, srcDir, mode)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ImportDir(t testing.TB, dir string, mode build.ImportMode) *build.Package {
	t.Helper()
	r0, err := build.ImportDir(dir, mode)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ArchChar(t testing.TB, goarch string) string {
	t.Helper()
	r0, err := build.ArchChar(goarch)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
