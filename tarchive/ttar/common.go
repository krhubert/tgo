package ttar

import (
	"archive/tar"
	"io/fs"
	"testing"
)

func FileInfoHeader(t testing.TB, fi fs.FileInfo, link string) *tar.Header {
	t.Helper()
	r0, err := tar.FileInfoHeader(fi, link)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
