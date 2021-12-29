package tdwarf

import (
	"debug/dwarf"
	"testing"
)

func New(t testing.TB, abbrev, aranges, frame, info, line, pubnames, ranges, str []byte) *dwarf.Data {
	t.Helper()
	r0, err := dwarf.New(abbrev, aranges, frame, info, line, pubnames, ranges, str)
	if err != nil {
		t.Fatal(err)
	}
	return r0
}
