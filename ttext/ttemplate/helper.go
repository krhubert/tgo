package ttemplate

import (
	"io/fs"
	"testing"
	"text/template"
)

func ParseFiles(t testing.TB, filenames ...string) *template.Template {
	t.Helper()
	r0, err := template.ParseFiles(filenames...)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseGlob(t testing.TB, pattern string) *template.Template {
	t.Helper()
	r0, err := template.ParseGlob(pattern)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}

func ParseFS(t testing.TB, fsys fs.FS, patterns ...string) *template.Template {
	t.Helper()
	r0, err := template.ParseFS(fsys, patterns...)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
