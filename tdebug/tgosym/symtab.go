package tgosym

import (
	"debug/gosym"
	"testing"
)

func NewTable(t testing.TB, symtab []byte, pcln *gosym.LineTable) *gosym.Table {
	t.Helper()
	r0, err := gosym.NewTable(symtab, pcln)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
