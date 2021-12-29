package tzip

import (
	"archive/zip"
	"io/fs"
	"testing"
)

func FileInfoHeader(t testing.TB, fi fs.FileInfo) *zip.FileHeader {
	t.Helper()
	r0, err := zip.FileInfoHeader(fi)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
