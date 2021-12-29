package tplugin

import (
	"plugin"
	"testing"
)

func Open(t testing.TB, path string) *plugin.Plugin {
	t.Helper()
	r0, err := plugin.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
